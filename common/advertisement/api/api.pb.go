// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// 广告轮播图
type AdsSwiper struct {
	// 广告轮播图
	Banner               []*SwiperItem `protobuf:"bytes,1,rep,name=Banner,proto3" json:"banner" form:"Banner"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AdsSwiper) Reset()         { *m = AdsSwiper{} }
func (m *AdsSwiper) String() string { return proto.CompactTextString(m) }
func (*AdsSwiper) ProtoMessage()    {}
func (*AdsSwiper) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *AdsSwiper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdsSwiper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdsSwiper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdsSwiper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdsSwiper.Merge(m, src)
}
func (m *AdsSwiper) XXX_Size() int {
	return m.Size()
}
func (m *AdsSwiper) XXX_DiscardUnknown() {
	xxx_messageInfo_AdsSwiper.DiscardUnknown(m)
}

var xxx_messageInfo_AdsSwiper proto.InternalMessageInfo

type SwiperItem struct {
	// 图片地址
	ImageUrl             string   `protobuf:"bytes,2,opt,name=ImageUrl,proto3" json:"imageUrl" form:"imageUrl"  gorm:"column:image_url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwiperItem) Reset()         { *m = SwiperItem{} }
func (m *SwiperItem) String() string { return proto.CompactTextString(m) }
func (*SwiperItem) ProtoMessage()    {}
func (*SwiperItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *SwiperItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwiperItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwiperItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwiperItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwiperItem.Merge(m, src)
}
func (m *SwiperItem) XXX_Size() int {
	return m.Size()
}
func (m *SwiperItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SwiperItem.DiscardUnknown(m)
}

var xxx_messageInfo_SwiperItem proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdsSwiper)(nil), "advertisement.service.v1.AdsSwiper")
	proto.RegisterType((*SwiperItem)(nil), "advertisement.service.v1.SwiperItem")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0x3d, 0x4f, 0xe3, 0x30,
	0x18, 0x6e, 0x5a, 0xa9, 0x6a, 0x7d, 0xd7, 0x25, 0xc3, 0xa9, 0x1f, 0xa7, 0x24, 0x97, 0x63, 0xe8,
	0x82, 0x23, 0xca, 0x82, 0xba, 0x35, 0x12, 0x42, 0x5d, 0x0b, 0x0c, 0xb0, 0x20, 0xa7, 0x75, 0x5d,
	0x4b, 0xb1, 0x1d, 0x39, 0x4e, 0x11, 0x13, 0x12, 0x7f, 0x81, 0x85, 0x9f, 0xd4, 0x11, 0x89, 0x3d,
	0x82, 0xc2, 0xd4, 0xb1, 0xbf, 0x00, 0x61, 0xa7, 0x2d, 0x0c, 0xdd, 0xfc, 0x3e, 0x5f, 0x7e, 0xdf,
	0x07, 0xd4, 0x51, 0x42, 0x61, 0x22, 0x85, 0x12, 0x76, 0x13, 0x4d, 0xe6, 0x58, 0x2a, 0x9a, 0x62,
	0x86, 0xb9, 0x82, 0x29, 0x96, 0x73, 0x3a, 0xc6, 0x70, 0x7e, 0xd4, 0x3e, 0x24, 0x54, 0xcd, 0xb2,
	0x08, 0x8e, 0x05, 0x0b, 0x88, 0x20, 0x22, 0xd0, 0x86, 0x28, 0x9b, 0xea, 0x49, 0x0f, 0xfa, 0x65,
	0x82, 0xda, 0x1d, 0x22, 0x04, 0x89, 0xf1, 0x4e, 0x85, 0x59, 0xa2, 0xee, 0x0a, 0xf2, 0x6f, 0x41,
	0xa2, 0x84, 0x06, 0x88, 0x73, 0xa1, 0x90, 0xa2, 0x82, 0xa7, 0x86, 0xf5, 0xa7, 0xa0, 0x3e, 0x98,
	0xa4, 0xe7, 0xb7, 0x34, 0xc1, 0xd2, 0xbe, 0x02, 0xd5, 0x10, 0x71, 0x8e, 0x65, 0xd3, 0xf2, 0x2a,
	0xdd, 0x5f, 0xbd, 0x03, 0xb8, 0x6f, 0x43, 0x68, 0x1c, 0x43, 0x85, 0x59, 0xd8, 0x59, 0xe5, 0x6e,
	0x35, 0xd2, 0xbe, 0x75, 0xee, 0x36, 0xa6, 0x42, 0xb2, 0xbe, 0x6f, 0x72, 0xfc, 0x51, 0x11, 0xe8,
	0x47, 0x00, 0xec, 0x2c, 0xf6, 0x05, 0xa8, 0x0d, 0x19, 0x22, 0xf8, 0x52, 0xc6, 0xcd, 0xb2, 0x67,
	0x75, 0xeb, 0xe1, 0xc9, 0x2a, 0x77, 0x6b, 0xb4, 0xc0, 0xd6, 0xb9, 0xdb, 0x35, 0x31, 0x1b, 0xc4,
	0xf7, 0x3c, 0xa2, 0x81, 0xb1, 0x88, 0x33, 0xc6, 0xfb, 0x1a, 0xbf, 0xc9, 0x64, 0xec, 0x8f, 0xb6,
	0x49, 0xbd, 0x7b, 0xd0, 0x18, 0x7c, 0xdf, 0xd7, 0xe6, 0xe0, 0xf7, 0x19, 0x56, 0xbb, 0xfb, 0xfe,
	0x40, 0xd3, 0x05, 0xdc, 0x14, 0x05, 0x4f, 0xbf, 0x8a, 0x6a, 0xff, 0xdf, 0x7f, 0xe7, 0xd6, 0xec,
	0xff, 0x7b, 0x78, 0xf9, 0x78, 0x2c, 0x77, 0xec, 0x56, 0xf0, 0x43, 0x1c, 0xcc, 0x04, 0xc3, 0x46,
	0x12, 0xb6, 0x16, 0x6f, 0x4e, 0x69, 0xb1, 0x74, 0xac, 0xe7, 0xa5, 0x63, 0xbd, 0x2e, 0x1d, 0xeb,
	0xe9, 0xdd, 0x29, 0x5d, 0x57, 0x50, 0x42, 0xa3, 0xaa, 0xfe, 0xf2, 0xf8, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0xe9, 0x8d, 0x47, 0xff, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdvertisementClient is the client API for Advertisement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdvertisementClient interface {
	// 获取首页广告轮播图
	GetAdsSwiper(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AdsSwiper, error)
}

type advertisementClient struct {
	cc *grpc.ClientConn
}

func NewAdvertisementClient(cc *grpc.ClientConn) AdvertisementClient {
	return &advertisementClient{cc}
}

func (c *advertisementClient) GetAdsSwiper(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AdsSwiper, error) {
	out := new(AdsSwiper)
	err := c.cc.Invoke(ctx, "/advertisement.service.v1.Advertisement/GetAdsSwiper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertisementServer is the server API for Advertisement service.
type AdvertisementServer interface {
	// 获取首页广告轮播图
	GetAdsSwiper(context.Context, *empty.Empty) (*AdsSwiper, error)
}

// UnimplementedAdvertisementServer can be embedded to have forward compatible implementations.
type UnimplementedAdvertisementServer struct {
}

func (*UnimplementedAdvertisementServer) GetAdsSwiper(ctx context.Context, req *empty.Empty) (*AdsSwiper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdsSwiper not implemented")
}

func RegisterAdvertisementServer(s *grpc.Server, srv AdvertisementServer) {
	s.RegisterService(&_Advertisement_serviceDesc, srv)
}

func _Advertisement_GetAdsSwiper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertisementServer).GetAdsSwiper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertisement.service.v1.Advertisement/GetAdsSwiper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertisementServer).GetAdsSwiper(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Advertisement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertisement.service.v1.Advertisement",
	HandlerType: (*AdvertisementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdsSwiper",
			Handler:    _Advertisement_GetAdsSwiper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *AdsSwiper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdsSwiper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdsSwiper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Banner) > 0 {
		for iNdEx := len(m.Banner) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Banner[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *SwiperItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwiperItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwiperItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x12
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
func (m *AdsSwiper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Banner) > 0 {
		for _, e := range m.Banner {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SwiperItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ImageUrl)
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
func (m *AdsSwiper) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AdsSwiper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdsSwiper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Banner", wireType)
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
			m.Banner = append(m.Banner, &SwiperItem{})
			if err := m.Banner[len(m.Banner)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *SwiperItem) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SwiperItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwiperItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
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
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
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
