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

	proto "github.com/gogo/protobuf/proto"
	payload "github.com/vdaas/vald/apis/grpc/payload"
	codes "github.com/vdaas/vald/internal/net/grpc/codes"
	status "github.com/vdaas/vald/internal/net/grpc/status"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("apis/proto/gateway/vald/vald.proto", fileDescriptor_667acc189b85e014)
}

var fileDescriptor_667acc189b85e014 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x8b, 0x13, 0x31,
	0x14, 0xc0, 0x77, 0x64, 0x99, 0xc5, 0x6c, 0x71, 0x25, 0xfe, 0x59, 0x77, 0x84, 0x22, 0x73, 0x92,
	0x45, 0x26, 0xa2, 0x37, 0x41, 0xc4, 0x6e, 0xdd, 0xa5, 0xb0, 0x75, 0x65, 0x17, 0x17, 0xf1, 0x96,
	0xce, 0x84, 0xd9, 0xc8, 0x74, 0x12, 0x93, 0x4c, 0xb5, 0x88, 0x17, 0xbf, 0x80, 0x07, 0xbf, 0x94,
	0x47, 0xc1, 0x2f, 0x20, 0xc5, 0x0f, 0x22, 0xf3, 0x92, 0xca, 0xb4, 0x9d, 0x52, 0x98, 0xb9, 0xb4,
	0xe9, 0x4b, 0xde, 0xaf, 0xef, 0x97, 0xc0, 0x7b, 0x28, 0xa4, 0x92, 0x6b, 0x22, 0x95, 0x30, 0x82,
	0xa4, 0xd4, 0xb0, 0x4f, 0x74, 0x4a, 0x26, 0x34, 0x4b, 0xe0, 0x23, 0x82, 0x38, 0xde, 0x2e, 0xd7,
	0xc1, 0x83, 0xca, 0x49, 0x49, 0xa7, 0x99, 0xa0, 0xc9, 0xfc, 0xdb, 0x9e, 0x0b, 0x9e, 0xa7, 0xdc,
	0x5c, 0x15, 0xa3, 0x28, 0x16, 0x63, 0x92, 0x0a, 0x91, 0x66, 0x0c, 0x52, 0x56, 0x96, 0x84, 0x4a,
	0x4e, 0x68, 0x9e, 0x0b, 0x43, 0x0d, 0x17, 0xb9, 0xb6, 0xe9, 0x4f, 0xbe, 0xef, 0xa2, 0xed, 0x4b,
	0x9a, 0x25, 0xf8, 0x18, 0xf9, 0xaf, 0x3e, 0x73, 0x6d, 0x34, 0xc6, 0xd1, 0xfc, 0x1f, 0xce, 0x46,
	0x1f, 0x58, 0x6c, 0xa2, 0x41, 0x3f, 0xa8, 0x89, 0x85, 0xb7, 0xbf, 0xfd, 0xfe, 0xfb, 0xe3, 0xda,
	0x0d, 0xdc, 0x21, 0x0c, 0x12, 0xc9, 0x17, 0x9e, 0x7c, 0xc5, 0x67, 0xc8, 0xbf, 0x60, 0x54, 0xc5,
	0x57, 0x78, 0xff, 0x7f, 0x8e, 0x0d, 0x44, 0xe7, 0xec, 0x63, 0xc1, 0xb4, 0x09, 0xee, 0xad, 0x6e,
	0x68, 0x29, 0x72, 0xcd, 0x42, 0x0c, 0xc8, 0x4e, 0xb8, 0x43, 0x34, 0xec, 0x3c, 0xf3, 0x0e, 0xf1,
	0x3b, 0x84, 0xec, 0xb1, 0xde, 0x74, 0xd0, 0xc7, 0x07, 0xcb, 0xb9, 0x83, 0xfe, 0x66, 0xec, 0x1d,
	0xc0, 0xee, 0x85, 0xc8, 0x61, 0x09, 0x4f, 0x4a, 0xf2, 0x09, 0xea, 0x5c, 0x18, 0xc5, 0xe8, 0xb8,
	0x79, 0xc1, 0x5b, 0x0f, 0xbd, 0xc7, 0x1e, 0x1e, 0xa2, 0x9b, 0x55, 0x50, 0xf3, 0x42, 0x2d, 0xee,
	0x35, 0xf2, 0x07, 0xb9, 0x66, 0xca, 0xe0, 0xbb, 0xcb, 0xd7, 0x7e, 0xc9, 0x62, 0x23, 0x54, 0x85,
	0xe0, 0xe2, 0xa7, 0x22, 0x86, 0x67, 0xad, 0xdc, 0x20, 0x07, 0x44, 0xe9, 0x79, 0x3c, 0xf7, 0x6c,
	0x4c, 0xb5, 0x75, 0x1d, 0xa1, 0xdd, 0x61, 0x91, 0x19, 0xee, 0x30, 0xfb, 0xf5, 0x18, 0x1d, 0x1c,
	0xac, 0xe3, 0xe8, 0x70, 0xab, 0x94, 0x7b, 0x2b, 0x13, 0x6a, 0x58, 0x2b, 0xb9, 0x02, 0x10, 0x0b,
	0x72, 0x8d, 0xa9, 0x8b, 0x72, 0x0e, 0xd3, 0x42, 0xae, 0xf5, 0xcb, 0x15, 0x72, 0xe5, 0xe5, 0x1a,
	0x53, 0x97, 0xe5, 0x5a, 0xbc, 0xdc, 0x29, 0xf2, 0xcf, 0xd9, 0x58, 0x4c, 0x58, 0x6d, 0x87, 0x58,
	0x5f, 0x82, 0xeb, 0x13, 0x87, 0x1d, 0xa2, 0x20, 0xdd, 0xf6, 0x89, 0xde, 0x5c, 0xad, 0x11, 0xd3,
	0x6a, 0xbd, 0x70, 0x5a, 0x0e, 0x71, 0x6b, 0x15, 0xb1, 0x41, 0x69, 0x88, 0xae, 0x9f, 0x30, 0x63,
	0x37, 0x6a, 0x2b, 0x58, 0x73, 0xe1, 0x95, 0xde, 0x27, 0x20, 0x6e, 0x9d, 0x8e, 0xd0, 0x9e, 0x75,
	0x6a, 0x06, 0x05, 0xa9, 0xde, 0xe8, 0xe7, 0xac, 0xeb, 0xfd, 0x9a, 0x75, 0xbd, 0x3f, 0xb3, 0xae,
	0x87, 0xee, 0x0b, 0x95, 0x46, 0x93, 0x84, 0x52, 0x1d, 0xc1, 0x70, 0x70, 0xe3, 0x02, 0x7e, 0xf4,
	0x76, 0xca, 0xce, 0xfd, 0x52, 0xf2, 0x37, 0xde, 0xfb, 0x47, 0x95, 0x31, 0x00, 0xc7, 0xed, 0x40,
	0xb1, 0x03, 0x40, 0xc9, 0x78, 0x61, 0xce, 0x8c, 0x7c, 0x68, 0xfe, 0x4f, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x35, 0xb9, 0x7d, 0xd0, 0x89, 0x06, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/vald.Vald/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Search(ctx context.Context, in *payload.Search_Request, opts ...grpc.CallOption) (*payload.Search_Response, error) {
	out := new(payload.Search_Response)
	err := c.cc.Invoke(ctx, "/vald.Vald/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) SearchByID(ctx context.Context, in *payload.Search_IDRequest, opts ...grpc.CallOption) (*payload.Search_Response, error) {
	out := new(payload.Search_Response)
	err := c.cc.Invoke(ctx, "/vald.Vald/SearchByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamSearch(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[0], "/vald.Vald/StreamSearch", opts...)
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
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[1], "/vald.Vald/StreamSearchByID", opts...)
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
	err := c.cc.Invoke(ctx, "/vald.Vald/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamInsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamInsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[2], "/vald.Vald/StreamInsert", opts...)
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
	err := c.cc.Invoke(ctx, "/vald.Vald/MultiInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Update(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.Vald/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamUpdate(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[3], "/vald.Vald/StreamUpdate", opts...)
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
	err := c.cc.Invoke(ctx, "/vald.Vald/MultiUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Upsert(ctx context.Context, in *payload.Object_Vector, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.Vald/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamUpsert(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamUpsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[4], "/vald.Vald/StreamUpsert", opts...)
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
	err := c.cc.Invoke(ctx, "/vald.Vald/MultiUpsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) Remove(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Location, error) {
	out := new(payload.Object_Location)
	err := c.cc.Invoke(ctx, "/vald.Vald/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamRemove(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamRemoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[5], "/vald.Vald/StreamRemove", opts...)
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
	err := c.cc.Invoke(ctx, "/vald.Vald/MultiRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) GetObject(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_Vector, error) {
	out := new(payload.Object_Vector)
	err := c.cc.Invoke(ctx, "/vald.Vald/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valdClient) StreamGetObject(ctx context.Context, opts ...grpc.CallOption) (Vald_StreamGetObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Vald_serviceDesc.Streams[6], "/vald.Vald/StreamGetObject", opts...)
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
		FullMethod: "/vald.Vald/Exists",
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
		FullMethod: "/vald.Vald/Search",
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
		FullMethod: "/vald.Vald/SearchByID",
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
		FullMethod: "/vald.Vald/Insert",
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
		FullMethod: "/vald.Vald/MultiInsert",
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
		FullMethod: "/vald.Vald/Update",
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
		FullMethod: "/vald.Vald/MultiUpdate",
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
		FullMethod: "/vald.Vald/Upsert",
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
		FullMethod: "/vald.Vald/MultiUpsert",
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
		FullMethod: "/vald.Vald/Remove",
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
		FullMethod: "/vald.Vald/MultiRemove",
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
		FullMethod: "/vald.Vald/GetObject",
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
	ServiceName: "vald.Vald",
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
	Metadata: "apis/proto/gateway/vald/vald.proto",
}
