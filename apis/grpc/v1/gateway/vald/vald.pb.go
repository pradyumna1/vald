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

package vald

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	payload "github.com/vdaas/vald/apis/grpc/v1/payload"
	codes "github.com/vdaas/vald/internal/net/grpc/codes"
	status "github.com/vdaas/vald/internal/net/grpc/status"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = golang_proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("apis/proto/v1/gateway/vald/vald.proto", fileDescriptor_afca6feed49a0850)
}

func init() {
	golang_proto.RegisterFile("apis/proto/v1/gateway/vald/vald.proto", fileDescriptor_afca6feed49a0850)
}

var fileDescriptor_afca6feed49a0850 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x37, 0x22, 0x29, 0x8e, 0xc1, 0x95, 0xc1, 0x15, 0x8c, 0x4b, 0x0e, 0x11, 0x41, 0x16,
	0xcc, 0x58, 0x3d, 0x08, 0xde, 0x2c, 0x5d, 0x34, 0xd0, 0x45, 0x69, 0xdd, 0x3d, 0x78, 0x10, 0xa6,
	0xc9, 0x98, 0x1d, 0x49, 0x33, 0x63, 0x66, 0x12, 0x2d, 0xe2, 0x41, 0xbf, 0x82, 0x5f, 0xc8, 0xe3,
	0x1e, 0x05, 0xbf, 0x80, 0x74, 0xfd, 0x20, 0x92, 0x37, 0x29, 0xa4, 0x35, 0x5d, 0x85, 0xf4, 0xd2,
	0x4e, 0xe6, 0xbd, 0xf7, 0xcb, 0xfb, 0xe7, 0xc1, 0xfb, 0xa3, 0xbb, 0x54, 0x72, 0x45, 0x64, 0x2e,
	0xb4, 0x20, 0x65, 0x9f, 0x24, 0x54, 0xb3, 0x0f, 0x74, 0x4e, 0x4a, 0x9a, 0xc6, 0xf0, 0x13, 0x40,
	0x08, 0xf7, 0xe0, 0x5c, 0xf6, 0xdd, 0x3b, 0xab, 0xf9, 0x92, 0xce, 0x53, 0x41, 0xe3, 0xe5, 0xbf,
	0xc9, 0x76, 0xef, 0x27, 0x5c, 0x9f, 0x16, 0xd3, 0x20, 0x12, 0x33, 0x92, 0x88, 0x44, 0x98, 0xfc,
	0x69, 0xf1, 0x16, 0x9e, 0x4c, 0x71, 0x75, 0xaa, 0xd3, 0x1f, 0xaf, 0xa7, 0x27, 0x42, 0x24, 0x29,
	0x83, 0x37, 0x99, 0x23, 0xa1, 0x92, 0x13, 0x9a, 0x65, 0x42, 0x53, 0xcd, 0x45, 0xa6, 0x4c, 0xe1,
	0xc3, 0x2f, 0x0e, 0xba, 0x7c, 0x42, 0xd3, 0x18, 0x8f, 0x90, 0x7d, 0xf8, 0x91, 0x2b, 0xad, 0xf0,
	0x5e, 0xb0, 0x6c, 0xa5, 0xec, 0x07, 0x2f, 0xa6, 0xef, 0x58, 0xa4, 0x83, 0x70, 0xe8, 0xb6, 0x5f,
	0xfb, 0x37, 0xbe, 0xfe, 0xfc, 0xfd, 0xed, 0xd2, 0x35, 0xec, 0x10, 0x06, 0xe5, 0xe4, 0x13, 0x8f,
	0x3f, 0xe3, 0x63, 0x64, 0x4f, 0x18, 0xcd, 0xa3, 0x53, 0xec, 0x36, 0xcb, 0xcc, 0x5d, 0x30, 0x66,
	0xef, 0x0b, 0xa6, 0xb4, 0x7b, 0xbb, 0x35, 0xa6, 0xa4, 0xc8, 0x14, 0xf3, 0x31, 0x80, 0x1d, 0xbf,
	0x47, 0x14, 0x44, 0x9e, 0x58, 0x07, 0xf8, 0x0d, 0x42, 0x26, 0x6d, 0x30, 0x0f, 0x87, 0x78, 0xbf,
	0xa5, 0x3c, 0x1c, 0xfe, 0x17, 0x7c, 0x0f, 0xe0, 0xbb, 0x3e, 0xaa, 0xe1, 0x84, 0xc7, 0x15, 0xff,
	0x08, 0x39, 0x13, 0x9d, 0x33, 0x3a, 0xeb, 0xda, 0xfc, 0xce, 0x3d, 0xeb, 0x81, 0x85, 0x27, 0xe8,
	0x7a, 0x13, 0xd7, 0xb5, 0x69, 0x03, 0x7d, 0x85, 0xec, 0x30, 0x53, 0x2c, 0xd7, 0xf8, 0x56, 0xcb,
	0x44, 0x4e, 0x58, 0xa4, 0x45, 0xbe, 0xca, 0xa9, 0x43, 0x23, 0x11, 0xc1, 0xe8, 0x1b, 0x5f, 0x96,
	0x03, 0xa8, 0x52, 0x3e, 0x5a, 0x2a, 0xef, 0xc8, 0x36, 0x3d, 0x86, 0xe8, 0xea, 0x51, 0x91, 0x6a,
	0x5e, 0xc3, 0xdc, 0x8d, 0x30, 0xe5, 0xee, 0x5f, 0x40, 0x53, 0xfe, 0x4e, 0x25, 0xf7, 0x58, 0xc6,
	0x54, 0xb3, 0x2d, 0xc8, 0x2d, 0x00, 0xb4, 0x22, 0xb7, 0x23, 0x7b, 0x55, 0x6e, 0x0d, 0xeb, 0x2c,
	0x77, 0x4b, 0xd3, 0x2d, 0xe4, 0x5f, 0xd3, 0xed, 0xc8, 0x5e, 0x97, 0xdb, 0x79, 0xba, 0x63, 0x64,
	0x8f, 0xd9, 0x4c, 0x94, 0x6c, 0xd3, 0xd6, 0xb9, 0xb0, 0x9d, 0x7a, 0xf7, 0x1c, 0x38, 0x24, 0x07,
	0x88, 0xd9, 0x3d, 0xcf, 0x97, 0x62, 0x3b, 0x90, 0x8d, 0xd0, 0xc3, 0x5a, 0x68, 0x0d, 0xba, 0xd9,
	0x0a, 0xfa, 0xb7, 0xc8, 0x09, 0xba, 0xf2, 0x8c, 0x69, 0x13, 0xd8, 0xd4, 0xcd, 0xe6, 0x89, 0x34,
	0x36, 0xac, 0x80, 0x7b, 0xa3, 0x32, 0x44, 0xbb, 0x46, 0x65, 0x17, 0x34, 0xc8, 0x1c, 0xe8, 0xb3,
	0x85, 0x67, 0xfd, 0x58, 0x78, 0xd6, 0xaf, 0x85, 0x67, 0x7d, 0x3f, 0xf7, 0xac, 0xb3, 0x73, 0xcf,
	0x42, 0xbe, 0xc8, 0x93, 0xa0, 0x8c, 0x29, 0x55, 0x01, 0x38, 0x17, 0x95, 0xbc, 0x02, 0xd4, 0xf6,
	0x06, 0x77, 0x83, 0x5e, 0x65, 0x1d, 0x4f, 0x25, 0x7f, 0x69, 0xbd, 0x26, 0x0d, 0x07, 0x82, 0x2a,
	0x63, 0x80, 0xc6, 0x81, 0x72, 0x19, 0xad, 0x5b, 0xe3, 0xd4, 0x06, 0x03, 0x7a, 0xf4, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x3c, 0x96, 0x40, 0xd9, 0x3f, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValdClient is the client API for Vald service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValdClient interface {
	Exists(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_ID, error)
	Search(ctx context.Context, in *payload.Search_Request, opts ...grpc.CallOption) (*payload.Search_Response, error)
	SearchByID(ctx context.Context, in *payload.Search_IDRequest, opts ...grpc.CallOption) (*payload.Search_Response, error)
	StreamSearch(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamSearchClient, error)
	StreamSearchByID(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamSearchByIDClient, error)
	Insert(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error)
	StreamInsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamInsertClient, error)
	MultiInsert(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error)
	Update(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error)
	StreamUpdate(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpdateClient, error)
	MultiUpdate(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error)
	Upsert(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error)
	StreamUpsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpsertClient, error)
	MultiUpsert(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error)
	Remove(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Location, error)
	StreamRemove(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamRemoveClient, error)
	MultiRemove(ctx context.Context, in *payload.Object_IDs, opts ...grpc.CallOption) (*payload.Object_Locations, error)
	GetObject(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Vector, error)
	StreamGetObject(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamGetObjectClient, error)
}

type valdClient struct {
	cc *grpc.ClientConn
}

func NewValdClient(cc *grpc.ClientConn) ValdClient {
	return &valdClient{cc}
}

func (c *valdClient) Exists(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_ID, error) {
	out := new(payload.Object_ID)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Search(ctx context.Context, in *payload.Search_Request, opts ...grpc.CallOption) (*payload.Search_Response, error) {
	out := new(payload.Search_Response)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) SearchByID(ctx context.Context, in *payload.Search_IDRequest, opts ...grpc.CallOption) (*payload.Search_Response, error) {
	out := new(payload.Search_Response)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/SearchByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamSearch(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[0], "/vald.v1.Vald/StreamSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamSearchClient{stream}
	return x, nil
}

type Vald_StreamSearchClient interface {
	Send(*payload.Search_Request) error
	Recv() (*payload.Search_Response, error)
	grpc.ClientStream
}

type valdStreamSearchClient struct {
	grpc.ClientStream
}

func (x *valdStreamSearchClient) Send(m *payload.Search_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamSearchClient) Recv() (*payload.Search_Response, error) {
	m := new(payload.Search_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) StreamSearchByID(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamSearchByIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[1], "/vald.v1.Vald/StreamSearchByID", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamSearchByIDClient{stream}
	return x, nil
}

type Vald_StreamSearchByIDClient interface {
	Send(*payload.Search_IDRequest) error
	Recv() (*payload.Search_Response, error)
	grpc.ClientStream
}

type valdStreamSearchByIDClient struct {
	grpc.ClientStream
}

func (x *valdStreamSearchByIDClient) Send(m *payload.Search_IDRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamSearchByIDClient) Recv() (*payload.Search_Response, error) {
	m := new(payload.Search_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) Insert(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamInsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamInsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[2], "/vald.v1.Vald/StreamInsert", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamInsertClient{stream}
	return x, nil
}

type Vald_StreamInsertClient interface {
	Send(*payload.Object_Vector) error
	Recv() (*payload.Object_Location, error)
	grpc.ClientStream
}

type valdStreamInsertClient struct {
	grpc.ClientStream
}

func (x *valdStreamInsertClient) Send(m *payload.Object_Vector) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamInsertClient) Recv() (*payload.Object_Location, error) {
	m := new(payload.Object_Location)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) MultiInsert(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error) {
	out := new(payload.Object_Locations)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/MultiInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Update(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamUpdate(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[3], "/vald.v1.Vald/StreamUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamUpdateClient{stream}
	return x, nil
}

type Vald_StreamUpdateClient interface {
	Send(*payload.Object_Vector) error
	Recv() (*payload.Object_Location, error)
	grpc.ClientStream
}

type valdStreamUpdateClient struct {
	grpc.ClientStream
}

func (x *valdStreamUpdateClient) Send(m *payload.Object_Vector) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamUpdateClient) Recv() (*payload.Object_Location, error) {
	m := new(payload.Object_Location)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) MultiUpdate(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error) {
	out := new(payload.Object_Locations)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/MultiUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Upsert(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamUpsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[4], "/vald.v1.Vald/StreamUpsert", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamUpsertClient{stream}
	return x, nil
}

type Vald_StreamUpsertClient interface {
	Send(*payload.Object_Vector) error
	Recv() (*payload.Object_Location, error)
	grpc.ClientStream
}

type valdStreamUpsertClient struct {
	grpc.ClientStream
}

func (x *valdStreamUpsertClient) Send(m *payload.Object_Vector) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamUpsertClient) Recv() (*payload.Object_Location, error) {
	m := new(payload.Object_Location)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) MultiUpsert(ctx context.Context, in *payload.Object_Vectors, opts ...grpc.CallOption) (*payload.Object_Locations, error) {
	out := new(payload.Object_Locations)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/MultiUpsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Remove(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamRemove(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamRemoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[5], "/vald.v1.Vald/StreamRemove", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamRemoveClient{stream}
	return x, nil
}

type Vald_StreamRemoveClient interface {
	Send(*payload.Object_ID) error
	Recv() (*payload.Object_Location, error)
	grpc.ClientStream
}

type valdStreamRemoveClient struct {
	grpc.ClientStream
}

func (x *valdStreamRemoveClient) Send(m *payload.Object_ID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamRemoveClient) Recv() (*payload.Object_Location, error) {
	m := new(payload.Object_Location)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *valdClient) MultiRemove(ctx context.Context, in *payload.Object_IDs, opts ...grpc.CallOption) (*payload.Object_Locations, error) {
	out := new(payload.Object_Locations)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/MultiRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) GetObject(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Vector, error) {
	out := new(payload.Object_Vector)
	err := c.cc.Invoke(ctx, "/vald.v1.Vald/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamGetObject(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamGetObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[6], "/vald.v1.Vald/StreamGetObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &valdStreamGetObjectClient{stream}
	return x, nil
}

type Vald_StreamGetObjectClient interface {
	Send(*payload.Object_ID) error
	Recv() (*payload.Object_Vector, error)
	grpc.ClientStream
}

type valdStreamGetObjectClient struct {
	grpc.ClientStream
}

func (x *valdStreamGetObjectClient) Send(m *payload.Object_ID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *valdStreamGetObjectClient) Recv() (*payload.Object_Vector, error) {
	m := new(payload.Object_Vector)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ValdServer is the server API for Vald service.
type ValdServer interface {
	Exists(context.Context, *payload.Object_ID) (*payload.Object_ID, error)
	Search(context.Context, *payload.Search_Request) (*payload.Search_Response, error)
	SearchByID(context.Context, *payload.Search_IDRequest) (*payload.Search_Response, error)
	StreamSearch(Vald_StreamSearchServer) error
	StreamSearchByID(Vald_StreamSearchByIDServer) error
	Insert(context.Context, *payload.Object_Vector) (*payload.Object_Location, error)
	StreamInsert(Vald_StreamInsertServer) error
	MultiInsert(context.Context, *payload.Object_Vectors) (*payload.Object_Locations, error)
	Update(context.Context, *payload.Object_Vector) (*payload.Object_Location, error)
	StreamUpdate(Vald_StreamUpdateServer) error
	MultiUpdate(context.Context, *payload.Object_Vectors) (*payload.Object_Locations, error)
	Upsert(context.Context, *payload.Object_Vector) (*payload.Object_Location, error)
	StreamUpsert(Vald_StreamUpsertServer) error
	MultiUpsert(context.Context, *payload.Object_Vectors) (*payload.Object_Locations, error)
	Remove(context.Context, *payload.Object_ID) (*payload.Object_Location, error)
	StreamRemove(Vald_StreamRemoveServer) error
	MultiRemove(context.Context, *payload.Object_IDs) (*payload.Object_Locations, error)
	GetObject(context.Context, *payload.Object_ID) (*payload.Object_Vector, error)
	StreamGetObject(Vald_StreamGetObjectServer) error
}

// UnimplementedValdServer can be embedded to have forward compatible implementations.
type UnimplementedValdServer struct{}

func (*UnimplementedValdServer) Exists(ctx context.Context, req *payload.Object_ID) (*payload.Object_ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}

func (*UnimplementedValdServer) Search(ctx context.Context, req *payload.Search_Request) (*payload.Search_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func (*UnimplementedValdServer) SearchByID(ctx context.Context, req *payload.Search_IDRequest) (*payload.Search_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByID not implemented")
}

func (*UnimplementedValdServer) StreamSearch(srv Vald_StreamSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSearch not implemented")
}

func (*UnimplementedValdServer) StreamSearchByID(srv Vald_StreamSearchByIDServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSearchByID not implemented")
}

func (*UnimplementedValdServer) Insert(ctx context.Context, req *payload.Object_Vector) (*payload.Object_Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func (*UnimplementedValdServer) StreamInsert(srv Vald_StreamInsertServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamInsert not implemented")
}

func (*UnimplementedValdServer) MultiInsert(ctx context.Context, req *payload.Object_Vectors) (*payload.Object_Locations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiInsert not implemented")
}

func (*UnimplementedValdServer) Update(ctx context.Context, req *payload.Object_Vector) (*payload.Object_Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (*UnimplementedValdServer) StreamUpdate(srv Vald_StreamUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUpdate not implemented")
}

func (*UnimplementedValdServer) MultiUpdate(ctx context.Context, req *payload.Object_Vectors) (*payload.Object_Locations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiUpdate not implemented")
}

func (*UnimplementedValdServer) Upsert(ctx context.Context, req *payload.Object_Vector) (*payload.Object_Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}

func (*UnimplementedValdServer) StreamUpsert(srv Vald_StreamUpsertServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamUpsert not implemented")
}

func (*UnimplementedValdServer) MultiUpsert(ctx context.Context, req *payload.Object_Vectors) (*payload.Object_Locations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiUpsert not implemented")
}

func (*UnimplementedValdServer) Remove(ctx context.Context, req *payload.Object_ID) (*payload.Object_Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func (*UnimplementedValdServer) StreamRemove(srv Vald_StreamRemoveServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRemove not implemented")
}

func (*UnimplementedValdServer) MultiRemove(ctx context.Context, req *payload.Object_IDs) (*payload.Object_Locations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiRemove not implemented")
}

func (*UnimplementedValdServer) GetObject(ctx context.Context, req *payload.Object_ID) (*payload.Object_Vector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}

func (*UnimplementedValdServer) StreamGetObject(srv Vald_StreamGetObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamGetObject not implemented")
}

func RegisterValdServer(s *grpc.Server, srv ValdServer) {
	s.RegisterService(&_Vald_serviceDesc, srv)
}

func _Vald_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Exists(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Search_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Search(ctx, req.(*payload.Search_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_SearchByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Search_IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).SearchByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/SearchByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).SearchByID(ctx, req.(*payload.Search_IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamSearch(&valdStreamSearchServer{stream})
}

type Vald_StreamSearchServer interface {
	Send(*payload.Search_Response) error
	Recv() (*payload.Search_Request, error)
	grpc.ServerStream
}

type valdStreamSearchServer struct {
	grpc.ServerStream
}

func (x *valdStreamSearchServer) Send(m *payload.Search_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamSearchServer) Recv() (*payload.Search_Request, error) {
	m := new(payload.Search_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_StreamSearchByID_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamSearchByID(&valdStreamSearchByIDServer{stream})
}

type Vald_StreamSearchByIDServer interface {
	Send(*payload.Search_Response) error
	Recv() (*payload.Search_IDRequest, error)
	grpc.ServerStream
}

type valdStreamSearchByIDServer struct {
	grpc.ServerStream
}

func (x *valdStreamSearchByIDServer) Send(m *payload.Search_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamSearchByIDServer) Recv() (*payload.Search_IDRequest, error) {
	m := new(payload.Search_IDRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Insert(ctx, req.(*payload.Object_Vector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamInsert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamInsert(&valdStreamInsertServer{stream})
}

type Vald_StreamInsertServer interface {
	Send(*payload.Object_Location) error
	Recv() (*payload.Object_Vector, error)
	grpc.ServerStream
}

type valdStreamInsertServer struct {
	grpc.ServerStream
}

func (x *valdStreamInsertServer) Send(m *payload.Object_Location) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamInsertServer) Recv() (*payload.Object_Vector, error) {
	m := new(payload.Object_Vector)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_MultiInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).MultiInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/MultiInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).MultiInsert(ctx, req.(*payload.Object_Vectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Update(ctx, req.(*payload.Object_Vector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamUpdate(&valdStreamUpdateServer{stream})
}

type Vald_StreamUpdateServer interface {
	Send(*payload.Object_Location) error
	Recv() (*payload.Object_Vector, error)
	grpc.ServerStream
}

type valdStreamUpdateServer struct {
	grpc.ServerStream
}

func (x *valdStreamUpdateServer) Send(m *payload.Object_Location) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamUpdateServer) Recv() (*payload.Object_Vector, error) {
	m := new(payload.Object_Vector)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_MultiUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).MultiUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/MultiUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).MultiUpdate(ctx, req.(*payload.Object_Vectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Upsert(ctx, req.(*payload.Object_Vector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamUpsert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamUpsert(&valdStreamUpsertServer{stream})
}

type Vald_StreamUpsertServer interface {
	Send(*payload.Object_Location) error
	Recv() (*payload.Object_Vector, error)
	grpc.ServerStream
}

type valdStreamUpsertServer struct {
	grpc.ServerStream
}

func (x *valdStreamUpsertServer) Send(m *payload.Object_Location) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamUpsertServer) Recv() (*payload.Object_Vector, error) {
	m := new(payload.Object_Vector)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_MultiUpsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_Vectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).MultiUpsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/MultiUpsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).MultiUpsert(ctx, req.(*payload.Object_Vectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).Remove(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamRemove_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamRemove(&valdStreamRemoveServer{stream})
}

type Vald_StreamRemoveServer interface {
	Send(*payload.Object_Location) error
	Recv() (*payload.Object_ID, error)
	grpc.ServerStream
}

type valdStreamRemoveServer struct {
	grpc.ServerStream
}

func (x *valdStreamRemoveServer) Send(m *payload.Object_Location) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamRemoveServer) Recv() (*payload.Object_ID, error) {
	m := new(payload.Object_ID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Vald_MultiRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).MultiRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/MultiRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).MultiRemove(ctx, req.(*payload.Object_IDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValdServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vald.v1.Vald/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValdServer).GetObject(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vald_StreamGetObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValdServer).StreamGetObject(&valdStreamGetObjectServer{stream})
}

type Vald_StreamGetObjectServer interface {
	Send(*payload.Object_Vector) error
	Recv() (*payload.Object_ID, error)
	grpc.ServerStream
}

type valdStreamGetObjectServer struct {
	grpc.ServerStream
}

func (x *valdStreamGetObjectServer) Send(m *payload.Object_Vector) error {
	return x.ServerStream.SendMsg(m)
}

func (x *valdStreamGetObjectServer) Recv() (*payload.Object_ID, error) {
	m := new(payload.Object_ID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Vald_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vald.v1.Vald",
	HandlerType: (*ValdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exists",
			Handler:    _Vald_Exists_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Vald_Search_Handler,
		},
		{
			MethodName: "SearchByID",
			Handler:    _Vald_SearchByID_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _Vald_Insert_Handler,
		},
		{
			MethodName: "MultiInsert",
			Handler:    _Vald_MultiInsert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Vald_Update_Handler,
		},
		{
			MethodName: "MultiUpdate",
			Handler:    _Vald_MultiUpdate_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _Vald_Upsert_Handler,
		},
		{
			MethodName: "MultiUpsert",
			Handler:    _Vald_MultiUpsert_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Vald_Remove_Handler,
		},
		{
			MethodName: "MultiRemove",
			Handler:    _Vald_MultiRemove_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Vald_GetObject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSearch",
			Handler:       _Vald_StreamSearch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSearchByID",
			Handler:       _Vald_StreamSearchByID_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamInsert",
			Handler:       _Vald_StreamInsert_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamUpdate",
			Handler:       _Vald_StreamUpdate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamUpsert",
			Handler:       _Vald_StreamUpsert_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamRemove",
			Handler:       _Vald_StreamRemove_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamGetObject",
			Handler:       _Vald_StreamGetObject_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apis/proto/v1/gateway/vald/vald.proto",
}
