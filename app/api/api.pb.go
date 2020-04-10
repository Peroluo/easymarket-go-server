// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: easymarketgoserve/app/api/api.proto

package api

import (
	context "context"
	api "easymarketgoserve/common/goods/api"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("easymarketgoserve/app/api/api.proto", fileDescriptor_af0fe30529ecb96a)
}

var fileDescriptor_af0fe30529ecb96a = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x1b, 0x0b, 0x1e, 0x82, 0x14, 0xdc, 0xe2, 0xa1, 0xa9, 0x2c, 0x52, 0xaf, 0xba, 0x41,
	0x7d, 0x02, 0x45, 0xe8, 0xc5, 0x53, 0x8f, 0xde, 0xb6, 0x71, 0x5c, 0x97, 0x36, 0x99, 0x31, 0xbb,
	0x2d, 0xf4, 0xea, 0x2b, 0x78, 0xf1, 0x91, 0x7a, 0x14, 0x3c, 0x79, 0xd3, 0xe8, 0x83, 0xc8, 0x4e,
	0x22, 0xb6, 0x48, 0x0e, 0x81, 0x0c, 0xdf, 0xff, 0xcf, 0xff, 0xcf, 0xc6, 0xc7, 0xa0, 0xdd, 0x2a,
	0xd7, 0xe5, 0x0c, 0xbc, 0x41, 0x07, 0xe5, 0x12, 0x52, 0x4d, 0x94, 0x6a, 0xb2, 0xe1, 0x53, 0x54,
	0xa2, 0x47, 0xd1, 0xd3, 0x44, 0x2a, 0x40, 0x9b, 0x81, 0x5a, 0x9e, 0x25, 0xa7, 0xc6, 0xfa, 0x87,
	0xc5, 0x54, 0x65, 0x98, 0xa7, 0x06, 0x0d, 0xa6, 0x2c, 0x9b, 0x2e, 0xee, 0x79, 0xe2, 0x81, 0xff,
	0x6a, 0x7b, 0x32, 0x34, 0x88, 0x66, 0x0e, 0x7f, 0x2a, 0xc8, 0xc9, 0xaf, 0x1a, 0x78, 0xd8, 0x40,
	0x4e, 0x2c, 0x0a, 0xf4, 0xda, 0x5b, 0x2c, 0x5c, 0x43, 0x4f, 0xfe, 0xd7, 0xcb, 0x30, 0xcf, 0xb1,
	0x48, 0x0d, 0xe2, 0x9d, 0xdb, 0xee, 0x79, 0xfe, 0x1e, 0xc5, 0xdd, 0x4b, 0x22, 0x01, 0xf1, 0xde,
	0x18, 0xfc, 0x38, 0x28, 0x6e, 0xac, 0xf3, 0x22, 0x51, 0xac, 0xde, 0x38, 0x41, 0x31, 0x9c, 0xc0,
	0x63, 0x22, 0x5b, 0x58, 0x30, 0x4e, 0xc0, 0x8d, 0x06, 0x4f, 0x6f, 0xdf, 0xcf, 0x3b, 0x7d, 0xb1,
	0xcf, 0xef, 0x62, 0x36, 0xd7, 0xce, 0xe2, 0xde, 0x6f, 0xcc, 0x35, 0x78, 0x6d, 0xe7, 0xe2, 0xa8,
	0x65, 0x59, 0x8d, 0x43, 0x5c, 0x7b, 0x15, 0x37, 0x1a, 0x72, 0xd4, 0x81, 0xe8, 0x6f, 0x45, 0xd5,
	0xde, 0xab, 0xc1, 0xfa, 0x53, 0x76, 0xd6, 0x95, 0x8c, 0x5e, 0x2b, 0x19, 0x7d, 0x54, 0x32, 0x7a,
	0xf9, 0x92, 0x9d, 0xdb, 0xae, 0x26, 0x3b, 0xdd, 0xe5, 0xeb, 0x2f, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0xb2, 0x70, 0xa3, 0xcc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	// 获取商品列表
	GetGoodsList(ctx context.Context, in *api.GoodsReq, opts ...grpc.CallOption) (*api.GoodsListRes, error)
	// 获取商品详情
	GetGoodsDetail(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*api.GoodsRes, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) GetGoodsList(ctx context.Context, in *api.GoodsReq, opts ...grpc.CallOption) (*api.GoodsListRes, error) {
	out := new(api.GoodsListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetGoodsDetail(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*api.GoodsRes, error) {
	out := new(api.GoodsRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	// 获取商品列表
	GetGoodsList(context.Context, *api.GoodsReq) (*api.GoodsListRes, error)
	// 获取商品详情
	GetGoodsDetail(context.Context, *api.GoodsDetailReq) (*api.GoodsRes, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) GetGoodsList(ctx context.Context, req *api.GoodsReq) (*api.GoodsListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}
func (*UnimplementedAppServer) GetGoodsDetail(ctx context.Context, req *api.GoodsDetailReq) (*api.GoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGoodsList(ctx, req.(*api.GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GoodsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGoodsDetail(ctx, req.(*api.GoodsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.service.v1.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsList",
			Handler:    _App_GetGoodsList_Handler,
		},
		{
			MethodName: "GetGoodsDetail",
			Handler:    _App_GetGoodsDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "easymarketgoserve/app/api/api.proto",
}
