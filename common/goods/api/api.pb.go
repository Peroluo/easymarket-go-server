// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: easymarketgoserve/common/goods/api/api.proto

package api

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// 商品详情Req
type GoodsDetailReq struct {
	// 商品id
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsDetailReq) Reset()         { *m = GoodsDetailReq{} }
func (m *GoodsDetailReq) String() string { return proto.CompactTextString(m) }
func (*GoodsDetailReq) ProtoMessage()    {}
func (*GoodsDetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a7aecac4e359b73, []int{0}
}
func (m *GoodsDetailReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsDetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsDetailReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsDetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsDetailReq.Merge(m, src)
}
func (m *GoodsDetailReq) XXX_Size() int {
	return m.Size()
}
func (m *GoodsDetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsDetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsDetailReq proto.InternalMessageInfo

// 商品列表Req
type GoodsReq struct {
	// 是否是热销 1:是、0:否
	IsHot int32 `protobuf:"varint,1,opt,name=isHot,proto3" json:"isHot,omitempty" form:"isHot"`
	// 页数
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	// 每页数量
	Size_                int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" form:"size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsReq) Reset()         { *m = GoodsReq{} }
func (m *GoodsReq) String() string { return proto.CompactTextString(m) }
func (*GoodsReq) ProtoMessage()    {}
func (*GoodsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a7aecac4e359b73, []int{1}
}
func (m *GoodsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsReq.Merge(m, src)
}
func (m *GoodsReq) XXX_Size() int {
	return m.Size()
}
func (m *GoodsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsReq proto.InternalMessageInfo

// 商品列表Res
type GoodsListRes struct {
	// 商品列表
	GoodsList []*GoodsRes `protobuf:"bytes,1,rep,name=goodsList,proto3" json:"goodsList,omitempty" form:"goodsList"`
	// 商品总数
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsListRes) Reset()         { *m = GoodsListRes{} }
func (m *GoodsListRes) String() string { return proto.CompactTextString(m) }
func (*GoodsListRes) ProtoMessage()    {}
func (*GoodsListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a7aecac4e359b73, []int{2}
}
func (m *GoodsListRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsListRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsListRes.Merge(m, src)
}
func (m *GoodsListRes) XXX_Size() int {
	return m.Size()
}
func (m *GoodsListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsListRes.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsListRes proto.InternalMessageInfo

// 商品详情Res
type GoodsRes struct {
	// 商品id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 商品名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 零食价格
	RetailPrice float64 `protobuf:"fixed64,3,opt,name=retail_price,json=retailPrice,proto3" json:"retail_price,omitempty"`
	// 商品简介
	GoodsBrief string `protobuf:"bytes,4,opt,name=goods_brief,json=goodsBrief,proto3" json:"goods_brief,omitempty"`
	// 商品图片
	ListPicUrl           string   `protobuf:"bytes,5,opt,name=list_pic_url,json=listPicUrl,proto3" json:"list_pic_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsRes) Reset()         { *m = GoodsRes{} }
func (m *GoodsRes) String() string { return proto.CompactTextString(m) }
func (*GoodsRes) ProtoMessage()    {}
func (*GoodsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a7aecac4e359b73, []int{3}
}
func (m *GoodsRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsRes.Merge(m, src)
}
func (m *GoodsRes) XXX_Size() int {
	return m.Size()
}
func (m *GoodsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsRes.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoodsDetailReq)(nil), "goods.service.v1.GoodsDetailReq")
	proto.RegisterType((*GoodsReq)(nil), "goods.service.v1.GoodsReq")
	proto.RegisterType((*GoodsListRes)(nil), "goods.service.v1.GoodsListRes")
	proto.RegisterType((*GoodsRes)(nil), "goods.service.v1.GoodsRes")
}

func init() {
	proto.RegisterFile("easymarketgoserve/common/goods/api/api.proto", fileDescriptor_0a7aecac4e359b73)
}

var fileDescriptor_0a7aecac4e359b73 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x9e, 0xbb, 0x16, 0x51, 0xb7, 0x94, 0xca, 0x30, 0x11, 0x32, 0x96, 0x14, 0x23, 0xa1, 0x1d,
	0x20, 0x11, 0xe3, 0xc6, 0xb1, 0x42, 0x1a, 0x07, 0x84, 0x26, 0x4b, 0x5c, 0xb8, 0x54, 0x69, 0xea,
	0x1a, 0x8b, 0x24, 0xce, 0x62, 0x77, 0xd2, 0x10, 0x27, 0x0e, 0xfc, 0x01, 0x38, 0xf0, 0x93, 0x76,
	0x44, 0xe2, 0x5e, 0xa0, 0xf0, 0x0b, 0xfa, 0x0b, 0x90, 0x9f, 0xb3, 0x76, 0x54, 0x2a, 0x87, 0x48,
	0xf1, 0xf7, 0x7d, 0x7e, 0xdf, 0x7b, 0xdf, 0x33, 0x7e, 0xc4, 0x13, 0x7d, 0x9e, 0x27, 0xd5, 0x3b,
	0x6e, 0x84, 0xd2, 0xbc, 0x3a, 0xe3, 0x71, 0xaa, 0xf2, 0x5c, 0x15, 0xb1, 0x50, 0x6a, 0xa2, 0xe3,
	0xa4, 0x94, 0xf6, 0x8b, 0xca, 0x4a, 0x19, 0x45, 0xfa, 0x00, 0x46, 0x56, 0x27, 0x53, 0x1e, 0x9d,
	0x3d, 0xf1, 0x1f, 0x0b, 0x69, 0xde, 0xce, 0xc6, 0x51, 0xaa, 0xf2, 0x58, 0x28, 0xa1, 0x62, 0x10,
	0x8e, 0x67, 0x53, 0x38, 0xc1, 0x01, 0xfe, 0x5c, 0x01, 0x7f, 0x5f, 0x28, 0x25, 0x32, 0xbe, 0x56,
	0xf1, 0xbc, 0x34, 0xe7, 0x35, 0x79, 0xaf, 0x26, 0xc1, 0xb3, 0x28, 0x94, 0x49, 0x8c, 0x54, 0x85,
	0x76, 0x2c, 0x8d, 0x71, 0xef, 0xd8, 0xba, 0x3f, 0xe7, 0x26, 0x91, 0x19, 0xe3, 0xa7, 0xe4, 0x00,
	0x37, 0xe4, 0xc4, 0x43, 0x03, 0x74, 0xd8, 0x1a, 0xde, 0x58, 0xce, 0xc3, 0xf6, 0x54, 0x55, 0xf9,
	0x33, 0x2a, 0x27, 0x94, 0x35, 0xe4, 0x84, 0x7e, 0xc0, 0xd7, 0xe1, 0x82, 0x95, 0x3e, 0xc4, 0x2d,
	0xa9, 0x5f, 0x28, 0x53, 0xab, 0xfb, 0xcb, 0x79, 0xd8, 0xad, 0xd5, 0x16, 0xa6, 0xcc, 0xd1, 0xe4,
	0x01, 0x6e, 0x96, 0x89, 0xe0, 0x5e, 0x03, 0x64, 0x37, 0x97, 0xf3, 0xb0, 0xe3, 0x64, 0x16, 0xa5,
	0x0c, 0x48, 0x2b, 0xd2, 0xf2, 0x3d, 0xf7, 0x76, 0x37, 0x45, 0x16, 0xa5, 0x0c, 0x48, 0xfa, 0x09,
	0xe1, 0x2e, 0xd8, 0xbf, 0x94, 0xda, 0x30, 0xae, 0xc9, 0x2b, 0xdc, 0x16, 0x97, 0x67, 0x0f, 0x0d,
	0x76, 0x0f, 0x3b, 0x47, 0x7e, 0xb4, 0x99, 0x67, 0x54, 0x77, 0xac, 0x87, 0xb7, 0x97, 0xf3, 0xb0,
	0xef, 0xca, 0xae, 0xae, 0x51, 0xb6, 0x2e, 0x61, 0x47, 0x32, 0xca, 0x24, 0x59, 0xdd, 0xeb, 0x95,
	0x91, 0x00, 0xa6, 0xcc, 0xd1, 0xf4, 0x0b, 0x5a, 0xe5, 0xa0, 0x49, 0x6f, 0x1d, 0x99, 0xcd, 0x88,
	0x10, 0xdc, 0x2c, 0x92, 0xdc, 0xcd, 0xdb, 0x66, 0xf0, 0x4f, 0xee, 0xe3, 0x6e, 0x05, 0x19, 0x8f,
	0xca, 0x4a, 0xa6, 0x6e, 0x4c, 0xc4, 0x3a, 0x0e, 0x3b, 0xb1, 0x10, 0x09, 0x71, 0x07, 0x1a, 0x19,
	0x8d, 0x2b, 0xc9, 0xa7, 0x5e, 0x13, 0x6e, 0x63, 0x80, 0x86, 0x16, 0x21, 0x03, 0xdc, 0xcd, 0xa4,
	0x36, 0xa3, 0x52, 0xa6, 0xa3, 0x59, 0x95, 0x79, 0x2d, 0xa7, 0xb0, 0xd8, 0x89, 0x4c, 0x5f, 0x57,
	0xd9, 0xd1, 0x0f, 0x84, 0x5b, 0xd0, 0x16, 0xc9, 0x71, 0xef, 0x98, 0x9b, 0x2b, 0xbb, 0x25, 0x83,
	0x2d, 0xb9, 0xac, 0x56, 0xef, 0xff, 0x27, 0x39, 0x7a, 0xf0, 0xf1, 0xfb, 0x9f, 0xcf, 0x8d, 0x3b,
	0x64, 0xaf, 0x7e, 0xc2, 0xe2, 0xdf, 0xe2, 0x02, 0x77, 0x2f, 0xed, 0x20, 0xc7, 0xed, 0xa5, 0x4e,
	0xfd, 0x60, 0x0b, 0x57, 0xef, 0x94, 0xee, 0x83, 0xd5, 0x1e, 0xb9, 0xb5, 0x61, 0x65, 0xf9, 0xe1,
	0xdd, 0x8b, 0x5f, 0xc1, 0xce, 0xc5, 0x22, 0x40, 0xdf, 0x16, 0x01, 0xfa, 0xb9, 0x08, 0xd0, 0xd7,
	0xdf, 0xc1, 0xce, 0x9b, 0xdd, 0xa4, 0x94, 0xe3, 0x6b, 0xf0, 0xa4, 0x9f, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0xbf, 0x86, 0xf1, 0x7e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoodsClient interface {
	// 获取商品详情
	GetGoodsDetail(ctx context.Context, in *GoodsDetailReq, opts ...grpc.CallOption) (*GoodsRes, error)
	// 获取商品列表
	GetGoodsList(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsListRes, error)
}

type goodsClient struct {
	cc *grpc.ClientConn
}

func NewGoodsClient(cc *grpc.ClientConn) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGoodsDetail(ctx context.Context, in *GoodsDetailReq, opts ...grpc.CallOption) (*GoodsRes, error) {
	out := new(GoodsRes)
	err := c.cc.Invoke(ctx, "/goods.service.v1.Goods/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsList(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsListRes, error) {
	out := new(GoodsListRes)
	err := c.cc.Invoke(ctx, "/goods.service.v1.Goods/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
type GoodsServer interface {
	// 获取商品详情
	GetGoodsDetail(context.Context, *GoodsDetailReq) (*GoodsRes, error)
	// 获取商品列表
	GetGoodsList(context.Context, *GoodsReq) (*GoodsListRes, error)
}

// UnimplementedGoodsServer can be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (*UnimplementedGoodsServer) GetGoodsDetail(ctx context.Context, req *GoodsDetailReq) (*GoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (*UnimplementedGoodsServer) GetGoodsList(ctx context.Context, req *GoodsReq) (*GoodsListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}

func RegisterGoodsServer(s *grpc.Server, srv GoodsServer) {
	s.RegisterService(&_Goods_serviceDesc, srv)
}

func _Goods_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.service.v1.Goods/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsDetail(ctx, req.(*GoodsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.service.v1.Goods/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsList(ctx, req.(*GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goods.service.v1.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsDetail",
			Handler:    _Goods_GetGoodsDetail_Handler,
		},
		{
			MethodName: "GetGoodsList",
			Handler:    _Goods_GetGoodsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "easymarketgoserve/common/goods/api/api.proto",
}

func (m *GoodsDetailReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsDetailReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsDetailReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GoodsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Size_ != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if m.Page != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if m.IsHot != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.IsHot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GoodsListRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsListRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsListRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Total != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.GoodsList) > 0 {
		for iNdEx := len(m.GoodsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoodsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GoodsRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListPicUrl) > 0 {
		i -= len(m.ListPicUrl)
		copy(dAtA[i:], m.ListPicUrl)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ListPicUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GoodsBrief) > 0 {
		i -= len(m.GoodsBrief)
		copy(dAtA[i:], m.GoodsBrief)
		i = encodeVarintApi(dAtA, i, uint64(len(m.GoodsBrief)))
		i--
		dAtA[i] = 0x22
	}
	if m.RetailPrice != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.RetailPrice))))
		i--
		dAtA[i] = 0x19
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoodsDetailReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApi(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoodsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsHot != 0 {
		n += 1 + sovApi(uint64(m.IsHot))
	}
	if m.Page != 0 {
		n += 1 + sovApi(uint64(m.Page))
	}
	if m.Size_ != 0 {
		n += 1 + sovApi(uint64(m.Size_))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoodsListRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GoodsList) > 0 {
		for _, e := range m.GoodsList {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovApi(uint64(m.Total))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoodsRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApi(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.RetailPrice != 0 {
		n += 9
	}
	l = len(m.GoodsBrief)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ListPicUrl)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoodsDetailReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoodsDetailReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsDetailReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoodsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoodsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsHot", wireType)
			}
			m.IsHot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsHot |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoodsListRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoodsListRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsListRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoodsList = append(m.GoodsList, &GoodsRes{})
			if err := m.GoodsList[len(m.GoodsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoodsRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoodsRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetailPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.RetailPrice = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsBrief", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoodsBrief = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListPicUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListPicUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
