//
// Copyright (C) 2019-2021 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package grpc provides grpc server logic
package grpc

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kpango/fuid"
	"github.com/vdaas/vald/apis/grpc/gateway/vald"
	"github.com/vdaas/vald/apis/grpc/payload"
	payloadv1 "github.com/vdaas/vald/apis/grpc/v1/payload"
	valdv1 "github.com/vdaas/vald/apis/grpc/v1/vald"
	"github.com/vdaas/vald/internal/core/algorithm"
	"github.com/vdaas/vald/internal/errgroup"
	"github.com/vdaas/vald/internal/errors"
	"github.com/vdaas/vald/internal/info"
	"github.com/vdaas/vald/internal/log"
	"github.com/vdaas/vald/internal/net/grpc"
	"github.com/vdaas/vald/internal/net/grpc/status"
	"github.com/vdaas/vald/internal/observability/trace"
	"github.com/vdaas/vald/internal/safety"
	"github.com/vdaas/vald/pkg/gateway/vald/service"
)

type server struct {
	eg                errgroup.Group
	gateway           service.Gateway
	metadata          service.Meta
	backup            service.Backup
	timeout           time.Duration
	filter            service.Filter
	replica           int
	streamConcurrency int
}

func New(opts ...Option) vald.ValdServer {
	s := new(server)

	for _, opt := range append(defaultOptions, opts...) {
		opt(s)
	}
	return s
}

func (s *server) Exists(ctx context.Context, meta *payload.Object_ID) (*payload.Object_ID, error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Exists")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	uuid, err := s.metadata.GetUUID(ctx, meta.GetId())
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("Exists API meta %s's uuid not found", meta.GetId()), err, meta.GetId(), info.Get())
	}
	return &payload.Object_ID{
		Id: uuid,
	}, nil
}

func (s *server) Search(ctx context.Context, req *payload.Search_Request) (res *payload.Search_Response, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Search")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	if len(req.Vector) < algorithm.MinimumVectorDimensionSize {
		return nil, errors.ErrInvalidDimensionSize(len(req.Vector), 0)
	}
	return s.search(ctx, req.GetConfig(),
		func(ctx context.Context, vc valdv1.Client, copts ...grpc.CallOption) (*payload.Search_Response, error) {
			res, err := vc.Search(ctx, &payloadv1.Search_Request{
				Vector: req.GetVector(),
			}, copts...)
			if err != nil {
				return nil, err
			}
			distances := make([]*payload.Object_Distance, 0, len(res.GetResults()))
			for _, r := range res.GetResults() {
				distances = append(distances, &payload.Object_Distance{
					Id:       r.GetId(),
					Distance: r.GetDistance(),
				})
			}
			return &payload.Search_Response{
				RequestId: res.GetRequestId(),
				Results:   distances,
			}, nil
		})
}

func (s *server) SearchByID(ctx context.Context, req *payload.Search_IDRequest) (
	res *payload.Search_Response, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.SearchByID")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	vec, err := s.GetObject(ctx, &payload.Object_ID{
		Id: req.GetId(),
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("SearchByID API meta %s's uuid not found", req.GetId()), err, req, info.Get())
	}
	return s.Search(ctx, &payload.Search_Request{
		Vector: vec.GetVector(),
		Config: req.GetConfig(),
	})
}

func (s *server) search(ctx context.Context, cfg *payload.Search_Config,
	f func(ctx context.Context, vc valdv1.Client, copts ...grpc.CallOption) (*payload.Search_Response, error)) (
	res *payload.Search_Response, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.search")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	var maxDist uint32
	atomic.StoreUint32(&maxDist, math.Float32bits(math.MaxFloat32))
	num := int(cfg.GetNum())
	res = new(payload.Search_Response)
	res.Results = make([]*payload.Object_Distance, 0, s.gateway.GetAgentCount(ctx)*num)
	dch := make(chan *payload.Object_Distance, cap(res.GetResults())/2)
	eg, ectx := errgroup.New(ctx)
	var cancel context.CancelFunc
	var timeout time.Duration
	if to := cfg.GetTimeout(); to != 0 {
		timeout = time.Duration(to)
	} else {
		timeout = s.timeout
	}
	ectx, cancel = context.WithTimeout(ectx, timeout)

	eg.Go(safety.RecoverFunc(func() error {
		defer cancel()
		visited := new(sync.Map)
		return s.gateway.BroadCast(ectx, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) error {
			r, err := f(ctx, vc, copts...)
			if err != nil {
				log.Debug("ignoring error:", err)
				return nil
			}
			for _, dist := range r.GetResults() {
				if dist == nil {
					continue
				}
				if dist.GetDistance() >= math.Float32frombits(atomic.LoadUint32(&maxDist)) {
					return nil
				}
				if _, already := visited.LoadOrStore(dist.GetId(), struct{}{}); !already {
					dch <- dist
				}
			}
			return nil
		})
	}))
	for {
		select {
		case <-ectx.Done():
			err = eg.Wait()
			if err != nil {
				log.Error("an error occurred while searching:", err)
			}
			close(dch)
			if len(res.GetResults()) > num && num != 0 {
				res.Results = res.Results[:num]
			}
			uuids := make([]string, 0, len(res.GetResults()))
			for _, r := range res.GetResults() {
				uuids = append(uuids, r.GetId())
			}
			if s.metadata != nil {
				metas, merr := s.metadata.GetMetas(ctx, uuids...)
				if merr != nil {
					log.Error("an error occurred during calling meta GetMetas:", merr)
					err = errors.Wrap(err, merr.Error())
				}
				for i, k := range metas {
					if len(k) != 0 {
						res.Results[i].Id = k
					}
				}
			}
			if s.filter != nil {
				r, ferr := s.filter.FilterSearch(ctx, res)
				if ferr == nil {
					res = r
				} else {
					err = errors.Wrap(err, ferr.Error())
				}
			}
			if err != nil {
				return res, status.WrapWithInternal(fmt.Sprintf("failed to search request %#v", cfg), err, info.Get())
			}
			return res, nil
		case dist := <-dch:
			rl := len(res.GetResults()) // result length
			if rl >= num && dist.GetDistance() >= math.Float32frombits(atomic.LoadUint32(&maxDist)) {
				continue
			}
			switch rl {
			case 0:
				res.Results = append(res.Results, dist)
			case 1:
				if res.GetResults()[0].GetDistance() <= dist.GetDistance() {
					res.Results = append(res.Results, dist)
				} else {
					res.Results = append([]*payload.Object_Distance{dist}, res.Results[0])
				}
			default:
				pos := rl
				for idx := rl; idx >= 1; idx-- {
					if res.GetResults()[idx-1].GetDistance() <= dist.GetDistance() {
						pos = idx - 1
						break
					}
				}

				switch {
				case pos == rl:
					res.Results = append([]*payload.Object_Distance{dist}, res.Results...)
				case pos == rl-1:
					res.Results = append(res.GetResults(), dist)
				case pos >= 0:
					res.Results = append(res.GetResults()[:pos+1], res.GetResults()[pos:]...)
					res.Results[pos+1] = dist
				}
			}
			rl = len(res.GetResults())
			if rl > num && num != 0 {
				res.Results = res.GetResults()[:num]
				rl = len(res.GetResults())
			}
			if distEnd := res.GetResults()[rl-1].GetDistance(); rl >= num &&
				distEnd < math.Float32frombits(atomic.LoadUint32(&maxDist)) {
				atomic.StoreUint32(&maxDist, math.Float32bits(distEnd))
			}
		}
	}
}

func (s *server) StreamSearch(stream vald.Vald_StreamSearchServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamSearch")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Search_Request) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.Search(ctx, data.(*payload.Search_Request))
		})
}

func (s *server) StreamSearchByID(stream vald.Vald_StreamSearchByIDServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamSearchByID")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Search_IDRequest) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.SearchByID(ctx, data.(*payload.Search_IDRequest))
		})
}

func (s *server) Insert(ctx context.Context, vec *payload.Object_Vector) (ce *payload.Object_Location, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Insert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
		err = errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
		if span != nil {
			span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
		}
		return nil, status.WrapWithInvalidArgument(fmt.Sprintf("Insert API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
	}
	meta := vec.GetId()
	exists, err := s.metadata.Exists(ctx, meta)
	if err != nil {
		log.Error("an error occurred while calling meta Exists:", err)
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(
			fmt.Sprintf("Insert API meta %s couldn't check meta already exists or not", meta), err, info.Get())
	}
	if exists {
		err = errors.Wrap(err, errors.ErrMetaDataAlreadyExists(meta).Error())
		if span != nil {
			span.SetStatus(trace.StatusCodeAlreadyExists(err.Error()))
		}
		return nil, status.WrapWithAlreadyExists(fmt.Sprintf("Insert API meta %s already exists", meta), err, info.Get())
	}

	uuid := fuid.String()
	err = s.metadata.SetUUIDandMeta(ctx, uuid, meta)
	if err != nil {
		log.Error("an error occurred during calling meta SetUUIDandMeta:", err)
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Insert API meta %s & uuid %s couldn't store", meta, uuid), err, info.Get())
	}
	vec.Id = uuid
	mu := new(sync.Mutex)
	targets := make([]string, 0, s.replica)
	err = s.gateway.DoMulti(ctx, s.replica, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) (err error) {
		_, err = vc.Insert(ctx, &payloadv1.Insert_Request{
			Vector: &payloadv1.Object_Vector{
				Id:     vec.GetId(),
				Vector: vec.GetVector(),
			},
		}, copts...)
		if err != nil {
			if err == errors.ErrRPCCallFailed(target, context.Canceled) {
				return nil
			}
			return err
		}
		target = strings.SplitN(target, ":", 2)[0]
		mu.Lock()
		targets = append(targets, target)
		mu.Unlock()
		return nil
	})
	if err != nil {
		err = errors.Wrapf(err, "Insert API (do multiple) failed to Insert uuid = %s\tmeta = %s\t info = %#v", uuid, meta, info.Get())
		log.Error(err)
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Insert API failed to Execute DoMulti error = %s", err.Error()), err, info.Get())
	}
	if s.backup != nil {
		vecs := &payloadv1.Backup_Vector{
			Uuid: uuid,
			Ips:  targets,
		}
		if vec != nil {
			vecs.Vector = vec.GetVector()
		}
		err = s.backup.Register(ctx, vecs)
		if err != nil {
			err = errors.Wrapf(err, "Insert API (backup.Register) failed to Backup Vectors = %#v\t info = %#v", vecs, info.Get())
			log.Error(err)
			if span != nil {
				span.SetStatus(trace.StatusCodeInternal(err.Error()))
			}
			return nil, status.WrapWithInternal(err.Error(), err)
		}
	}
	log.Debugf("Insert API insert succeeded to %v", targets)
	return new(payload.Object_Location), nil
}

func (s *server) StreamInsert(stream vald.Vald_StreamInsertServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamInsert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Object_Vector) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.Insert(ctx, data.(*payload.Object_Vector))
		})
}

func (s *server) MultiInsert(ctx context.Context, vecs *payload.Object_Vectors) (res *payload.Object_Locations, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.MultiInsert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	metaMap := make(map[string]string)
	metas := make([]string, 0, len(vecs.GetVectors()))
	reqs := make([]*payloadv1.Insert_Request, 0, len(vecs.GetVectors()))
	for _, vec := range vecs.GetVectors() {
		if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
			err = errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
			if span != nil {
				span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
			}
			return nil, status.WrapWithInvalidArgument(fmt.Sprintf("MultiInsert API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
		}
		uuid := fuid.String()
		meta := vec.GetId()
		metaMap[uuid] = meta
		metas = append(metas, meta)
		reqs = append(reqs, &payloadv1.Insert_Request{
			Vector: &payloadv1.Object_Vector{
				Vector: vec.GetVector(),
				Id:     uuid,
			},
		})
	}

	for _, meta := range metas {
		exists, err := s.metadata.Exists(ctx, meta)
		if err != nil {
			log.Error("an error occurred during calling meta Exists:", err)
			if span != nil {
				span.SetStatus(trace.StatusCodeInternal(err.Error()))
			}
			return nil, status.WrapWithInternal(
				fmt.Sprintf("MultiInsert API couldn't check metadata exists or not metas = %v", metas), err, info.Get())
		}
		if exists {
			if span != nil {
				span.SetStatus(trace.StatusCodeAlreadyExists(err.Error()))
			}
			return nil, status.WrapWithAlreadyExists(
				fmt.Sprintf("MultiInsert API failed metadata already exists meta = %s", meta), err, info.Get())
		}
	}

	err = s.metadata.SetUUIDandMetas(ctx, metaMap)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiInsert API failed SetUUIDandMetas %#v", metaMap), err, info.Get())
	}

	mu := new(sync.Mutex)
	targets := make([]string, 0, s.replica)
	gerr := s.gateway.DoMulti(ctx, s.replica, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) (err error) {
		_, err = vc.MultiInsert(ctx, &payloadv1.Insert_MultiRequest{
			Requests: reqs,
		}, copts...)
		if err != nil {
			return err
		}
		target = strings.SplitN(target, ":", 2)[0]
		mu.Lock()
		targets = append(targets, target)
		mu.Unlock()
		return nil
	})
	if gerr != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiInsert API failed request %#v", vecs), errors.Wrap(gerr, err.Error()), info.Get())
	}

	if s.backup != nil {
		mvecs := new(payloadv1.Backup_Vectors)
		mvecs.Vectors = make([]*payloadv1.Backup_Vector, 0, len(vecs.GetVectors()))
		for _, req := range reqs {
			vec := req.GetVector()
			uuid := vec.GetId()
			mvecs.Vectors = append(mvecs.Vectors, &payloadv1.Backup_Vector{
				Uuid:   uuid,
				Vector: vec.GetVector(),
				Ips:    targets,
			})
		}
		err = s.backup.RegisterMultiple(ctx, mvecs)
		if err != nil {
			if span != nil {
				span.SetStatus(trace.StatusCodeInternal(err.Error()))
			}
			return nil, status.WrapWithInternal(fmt.Sprintf("MultiInsert API failed RegisterMultiple %#v", mvecs), err, info.Get())
		}
	}
	return new(payload.Object_Locations), nil
}

func (s *server) Update(ctx context.Context, vec *payload.Object_Vector) (res *payload.Object_Location, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Update")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
		err = errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
		if span != nil {
			span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
		}
		return nil, status.WrapWithInvalidArgument(fmt.Sprintf("Update API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
	}
	meta := vec.GetId()
	uuid, err := s.metadata.GetUUID(ctx, meta)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("Update API failed GetUUID meta = %s", meta), err, info.Get())
	}
	vec.Id = uuid
	locs, err := s.backup.GetLocation(ctx, uuid)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("Update API failed GetLocation meta = %s, uuid = %s", meta, uuid), err, info.Get())
	}
	lmap := make(map[string]struct{}, len(locs))
	for _, loc := range locs {
		lmap[loc] = struct{}{}
	}
	err = s.gateway.BroadCast(ctx, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) error {
		target = strings.SplitN(target, ":", 2)[0]
		_, ok := lmap[target]
		if ok {
			_, err = vc.Update(ctx, &payloadv1.Update_Request{
				Vector: &payloadv1.Object_Vector{
					Vector: vec.GetVector(),
					Id:     vec.GetId(),
				},
			}, copts...)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Update API failed request %#v", vec), err, info.Get())
	}
	err = s.backup.Register(ctx, &payloadv1.Backup_Vector{
		Uuid:   uuid,
		Vector: vec.GetVector(),
		Ips:    locs,
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Update API failed backup %#v", vec), err, info.Get())
	}

	return new(payload.Object_Location), nil
}

func (s *server) StreamUpdate(stream vald.Vald_StreamUpdateServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamUpdate")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Object_Vector) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.Update(ctx, data.(*payload.Object_Vector))
		})
}

func (s *server) MultiUpdate(ctx context.Context, vecs *payload.Object_Vectors) (res *payload.Object_Locations, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.MultiUpdate")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	ids := make([]string, 0, len(vecs.GetVectors()))
	for _, vec := range vecs.GetVectors() {
		if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
			err = errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
			if span != nil {
				span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
			}
			return nil, status.WrapWithInvalidArgument(fmt.Sprintf("MultiUpdate API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
		}
		ids = append(ids, vec.GetId())
	}
	_, err = s.MultiRemove(ctx, &payload.Object_IDs{
		Ids: ids,
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiUpdate API failed Remove request %#v", ids), err, info.Get())
	}
	_, err = s.MultiInsert(ctx, vecs)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiUpdate API failed Insert request %#v", vecs), err, info.Get())
	}
	return new(payload.Object_Locations), nil
}

func (s *server) Upsert(ctx context.Context, vec *payload.Object_Vector) (*payload.Object_Location, error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Upsert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
		err := errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
		if span != nil {
			span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
		}
		return nil, status.WrapWithInvalidArgument(fmt.Sprintf("Upsert API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
	}
	meta := vec.GetId()
	exists, errs := s.metadata.Exists(ctx, meta)
	if errs != nil {
		log.Error("an error occurred during calling meta Exists:", errs)
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(errs.Error()))
		}
	}

	if exists {
		_, err := s.Update(ctx, vec)
		if err != nil {
			errs = errors.Wrap(errs, err.Error())
		}
	} else {
		_, err := s.Insert(ctx, vec)
		if err != nil {
			errs = errors.Wrap(errs, err.Error())
		}
	}

	return new(payload.Object_Location), errs
}

func (s *server) StreamUpsert(stream vald.Vald_StreamUpsertServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamUpsert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Object_Vector) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.Upsert(ctx, data.(*payload.Object_Vector))
		})
}

func (s *server) MultiUpsert(ctx context.Context, vecs *payload.Object_Vectors) (*payload.Object_Locations, error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.MultiUpsert")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	insertVecs := make([]*payload.Object_Vector, 0, len(vecs.GetVectors()))
	updateVecs := make([]*payload.Object_Vector, 0, len(vecs.GetVectors()))

	var errs error
	for _, vec := range vecs.GetVectors() {
		if len(vec.GetVector()) < algorithm.MinimumVectorDimensionSize {
			err := errors.ErrInvalidDimensionSize(len(vec.GetVector()), 0)
			if span != nil {
				span.SetStatus(trace.StatusCodeInvalidArgument(err.Error()))
			}
			return nil, status.WrapWithInvalidArgument(fmt.Sprintf("MultiUpsert API meta datas %v's vector invalid", vec.GetId()), err, info.Get())
		}
		exists, err := s.metadata.Exists(ctx, vec.GetId())
		if err != nil {
			log.Error("an error occurred during calling meta Exists:", err)
			if span != nil {
				span.SetStatus(trace.StatusCodeInternal(err.Error()))
			}
			errs = errors.Wrap(errs, err.Error())
		}

		if exists {
			updateVecs = append(updateVecs, vec)
		} else {
			insertVecs = append(insertVecs, vec)
		}
	}

	eg, ectx := errgroup.New(ctx)

	eg.Go(safety.RecoverFunc(func() error {
		var err error
		if len(updateVecs) > 0 {
			_, err = s.MultiUpdate(ectx, &payload.Object_Vectors{
				Vectors: updateVecs,
			})
		}
		return err
	}))

	eg.Go(safety.RecoverFunc(func() error {
		var err error
		if len(insertVecs) > 0 {
			_, err = s.MultiInsert(ectx, &payload.Object_Vectors{
				Vectors: insertVecs,
			})
		}
		return err
	}))

	err := eg.Wait()
	if err != nil {
		errs = errors.Wrap(errs, err.Error())
		return nil, status.WrapWithInternal("MultiUpsert API failed", errs, info.Get())
	}

	return new(payload.Object_Locations), errs
}

func (s *server) Remove(ctx context.Context, id *payload.Object_ID) (*payload.Object_Location, error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.Remove")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	meta := id.GetId()
	uuid, err := s.metadata.GetUUID(ctx, meta)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("Remove API meta %s's uuid not found", meta), err, info.Get())
	}
	locs, err := s.backup.GetLocation(ctx, uuid)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("Remove API failed GetLocation meta = %s, uuid = %s", meta, uuid), err, info.Get())
	}
	lmap := make(map[string]struct{}, len(locs))
	for _, loc := range locs {
		lmap[loc] = struct{}{}
	}
	err = s.gateway.BroadCast(ctx, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) error {
		_, ok := lmap[target]
		if ok {
			_, err = vc.Remove(ctx, &payloadv1.Remove_Request{
				Id: &payloadv1.Object_ID{
					Id: uuid,
				},
			}, copts...)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Remove API failed request uuid %s", uuid), err, info.Get())
	}
	_, err = s.metadata.DeleteMeta(ctx, uuid)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Remove API failed Delete metadata uuid = %s", uuid), err, info.Get())
	}
	err = s.backup.Remove(ctx, uuid)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("Remove API failed to Remove backup uuid = %s", uuid), err, info.Get())
	}
	return new(payload.Object_Location), nil
}

func (s *server) StreamRemove(stream vald.Vald_StreamRemoveServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamRemove")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Object_ID) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.Remove(ctx, data.(*payload.Object_ID))
		})
}

func (s *server) MultiRemove(ctx context.Context, ids *payload.Object_IDs) (res *payload.Object_Locations, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.MultiRemove")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	uuids, err := s.metadata.GetUUIDs(ctx, ids.GetIds()...)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("MultiRemove API meta datas %v's uuid not found", ids.GetIds()), err, info.Get())
	}
	lmap := make(map[string][]string, s.gateway.GetAgentCount(ctx))
	for _, uuid := range uuids {
		locs, err := s.backup.GetLocation(ctx, uuid)
		if err != nil {
			return nil, status.WrapWithNotFound(fmt.Sprintf("MultiRemove API failed to get uuid %s's location ", uuid), err, info.Get())
		}
		for _, loc := range locs {
			lmap[loc] = append(lmap[loc], uuid)
		}
	}
	err = s.gateway.BroadCast(ctx, func(ctx context.Context, target string, vc valdv1.Client, copts ...grpc.CallOption) error {
		uuids, ok := lmap[target]
		if ok {
			reqs := make([]*payloadv1.Remove_Request, 0, len(uuids))
			for _, uuid := range uuids {
				reqs = append(reqs, &payloadv1.Remove_Request{
					Id: &payloadv1.Object_ID{
						Id: uuid,
					},
				})
			}
			_, err := vc.MultiRemove(ctx, &payloadv1.Remove_MultiRequest{
				Requests: reqs,
			}, copts...)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiRemove API failed to request uuids %v metas %v ", uuids, ids.GetIds()), err, info.Get())
	}
	_, err = s.metadata.DeleteMetas(ctx, uuids...)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiRemove API failed to DeleteMetas uuids %v ", uuids), err, info.Get())
	}
	err = s.backup.RemoveMultiple(ctx, uuids...)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeInternal(err.Error()))
		}
		return nil, status.WrapWithInternal(fmt.Sprintf("MultiRemove API failed to Remove backup uuids %v ", uuids), err, info.Get())
	}
	return new(payload.Object_Locations), nil
}

func (s *server) GetObject(ctx context.Context, id *payload.Object_ID) (vec *payload.Object_Vector, err error) {
	ctx, span := trace.StartSpan(ctx, "vald/gateway-vald.GetObject")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	meta := id.GetId()
	uuid, err := s.metadata.GetUUID(ctx, meta)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("GetObject API meta %s's uuid not found", meta), err, info.Get())
	}
	mvec, err := s.backup.GetObject(ctx, uuid)
	if err != nil {
		if span != nil {
			span.SetStatus(trace.StatusCodeNotFound(err.Error()))
		}
		return nil, status.WrapWithNotFound(fmt.Sprintf("GetObject API meta %s uuid %s Object not found", meta, uuid), err, info.Get())
	}
	vec = &payload.Object_Vector{
		Id:     mvec.GetUuid(),
		Vector: mvec.GetVector(),
	}
	return vec, nil
}

func (s *server) StreamGetObject(stream vald.Vald_StreamGetObjectServer) error {
	ctx, span := trace.StartSpan(stream.Context(), "vald/gateway-vald.StreamGetObject")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	return grpc.BidirectionalStream(ctx, stream, s.streamConcurrency,
		func() interface{} { return new(payload.Object_ID) },
		func(ctx context.Context, data interface{}) (interface{}, error) {
			return s.GetObject(ctx, data.(*payload.Object_ID))
		})
}
