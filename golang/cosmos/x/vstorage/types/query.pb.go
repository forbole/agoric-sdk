// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agoric/vstorage/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryDataRequest is the vstorage path data query.
type QueryDataRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path" yaml:"path"`
}

func (m *QueryDataRequest) Reset()         { *m = QueryDataRequest{} }
func (m *QueryDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDataRequest) ProtoMessage()    {}
func (*QueryDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a26d6d1a170e94ae, []int{0}
}
func (m *QueryDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDataRequest.Merge(m, src)
}
func (m *QueryDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDataRequest proto.InternalMessageInfo

func (m *QueryDataRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// QueryDataResponse is the vstorage path data response.
type QueryDataResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value" yaml:"value"`
}

func (m *QueryDataResponse) Reset()         { *m = QueryDataResponse{} }
func (m *QueryDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDataResponse) ProtoMessage()    {}
func (*QueryDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a26d6d1a170e94ae, []int{1}
}
func (m *QueryDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDataResponse.Merge(m, src)
}
func (m *QueryDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDataResponse proto.InternalMessageInfo

func (m *QueryDataResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// QueryChildrenRequest is the vstorage path children query.
type QueryChildrenRequest struct {
	Path       string             `protobuf:"bytes,1,opt,name=path,proto3" json:"path" yaml:"path"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChildrenRequest) Reset()         { *m = QueryChildrenRequest{} }
func (m *QueryChildrenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChildrenRequest) ProtoMessage()    {}
func (*QueryChildrenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a26d6d1a170e94ae, []int{2}
}
func (m *QueryChildrenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChildrenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChildrenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChildrenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChildrenRequest.Merge(m, src)
}
func (m *QueryChildrenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChildrenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChildrenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChildrenRequest proto.InternalMessageInfo

func (m *QueryChildrenRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *QueryChildrenRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryChildrenResponse is the vstorage path children response.
type QueryChildrenResponse struct {
	Children   []string            `protobuf:"bytes,1,rep,name=children,proto3" json:"children" yaml:"children"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChildrenResponse) Reset()         { *m = QueryChildrenResponse{} }
func (m *QueryChildrenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChildrenResponse) ProtoMessage()    {}
func (*QueryChildrenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a26d6d1a170e94ae, []int{3}
}
func (m *QueryChildrenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChildrenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChildrenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChildrenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChildrenResponse.Merge(m, src)
}
func (m *QueryChildrenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChildrenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChildrenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChildrenResponse proto.InternalMessageInfo

func (m *QueryChildrenResponse) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *QueryChildrenResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDataRequest)(nil), "agoric.vstorage.QueryDataRequest")
	proto.RegisterType((*QueryDataResponse)(nil), "agoric.vstorage.QueryDataResponse")
	proto.RegisterType((*QueryChildrenRequest)(nil), "agoric.vstorage.QueryChildrenRequest")
	proto.RegisterType((*QueryChildrenResponse)(nil), "agoric.vstorage.QueryChildrenResponse")
}

func init() { proto.RegisterFile("agoric/vstorage/query.proto", fileDescriptor_a26d6d1a170e94ae) }

var fileDescriptor_a26d6d1a170e94ae = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xb1, 0x6f, 0x13, 0x3f,
	0x14, 0x8e, 0xf3, 0x6b, 0x7f, 0x6a, 0x5d, 0xa4, 0x82, 0x55, 0x44, 0x08, 0xd5, 0x5d, 0xb0, 0xa0,
	0x44, 0x20, 0x6c, 0xb5, 0x6c, 0x74, 0x40, 0x84, 0x0a, 0x56, 0x38, 0x89, 0x85, 0xed, 0x25, 0xb1,
	0x9c, 0x13, 0x97, 0xf3, 0xf5, 0xec, 0x44, 0x44, 0x88, 0xa5, 0x8c, 0x2c, 0x48, 0xcc, 0xfc, 0x3f,
	0x8c, 0x95, 0x58, 0x98, 0x4e, 0x28, 0x61, 0xca, 0x98, 0xbf, 0x00, 0x9d, 0xed, 0xb4, 0xd1, 0x51,
	0x51, 0x89, 0xed, 0xde, 0xf7, 0xbd, 0xf7, 0xbd, 0xef, 0x3e, 0xdb, 0xf8, 0x16, 0x48, 0x95, 0xc7,
	0x3d, 0x3e, 0xd6, 0x46, 0xe5, 0x20, 0x05, 0x3f, 0x1e, 0x89, 0x7c, 0xc2, 0xb2, 0x5c, 0x19, 0x45,
	0xb6, 0x1d, 0xc9, 0x96, 0x64, 0x73, 0x47, 0x2a, 0xa9, 0x2c, 0xc7, 0xcb, 0x2f, 0xd7, 0xd6, 0xbc,
	0xdf, 0x53, 0x7a, 0xa8, 0x34, 0xef, 0x82, 0xf6, 0xf3, 0x7c, 0xbc, 0xdf, 0x15, 0x06, 0xf6, 0x79,
	0x06, 0x32, 0x4e, 0xc1, 0xc4, 0x2a, 0xf5, 0xbd, 0xbb, 0x52, 0x29, 0x99, 0x08, 0x0e, 0x59, 0xcc,
	0x21, 0x4d, 0x95, 0xb1, 0xa4, 0x76, 0x2c, 0x7d, 0x82, 0xaf, 0xbe, 0x2a, 0xe7, 0x8f, 0xc0, 0x40,
	0x24, 0x8e, 0x47, 0x42, 0x1b, 0xf2, 0x00, 0xaf, 0x65, 0x60, 0x06, 0x0d, 0xd4, 0x42, 0xed, 0xcd,
	0xce, 0x8d, 0x79, 0x11, 0xda, 0x7a, 0x51, 0x84, 0x5b, 0x13, 0x18, 0x26, 0x8f, 0x69, 0x59, 0xd1,
	0xc8, 0x82, 0xf4, 0x08, 0x5f, 0x5b, 0x11, 0xd0, 0x99, 0x4a, 0xb5, 0x20, 0x1c, 0xaf, 0x8f, 0x21,
	0x19, 0x09, 0x2f, 0x71, 0x73, 0x5e, 0x84, 0x0e, 0x58, 0x14, 0xe1, 0x15, 0xa7, 0x61, 0x4b, 0x1a,
	0x39, 0x98, 0x7e, 0x42, 0x78, 0xc7, 0xca, 0x3c, 0x1b, 0xc4, 0x49, 0x3f, 0x17, 0xe9, 0xbf, 0x78,
	0x21, 0xcf, 0x31, 0x3e, 0xff, 0xfd, 0x46, 0xbd, 0x85, 0xda, 0x5b, 0x07, 0x7b, 0xcc, 0x65, 0xc5,
	0xca, 0xac, 0x98, 0xcb, 0xda, 0x67, 0xc5, 0x5e, 0x82, 0x14, 0x7e, 0x51, 0xb4, 0x32, 0x49, 0xbf,
	0x22, 0x7c, 0xbd, 0xe2, 0xc6, 0xff, 0xd8, 0x21, 0xde, 0xe8, 0x79, 0xac, 0x81, 0x5a, 0xff, 0xb5,
	0x37, 0x3b, 0xe1, 0xbc, 0x08, 0xcf, 0xb0, 0x45, 0x11, 0x6e, 0x3b, 0x5b, 0x4b, 0x84, 0x46, 0x67,
	0x24, 0x79, 0x71, 0x81, 0xbd, 0x7b, 0x97, 0xda, 0x73, 0x9b, 0x57, 0xfd, 0x1d, 0x9c, 0xd4, 0xf1,
	0xba, 0xf5, 0x47, 0x34, 0x5e, 0x2b, 0x83, 0x27, 0xb7, 0x59, 0xe5, 0xe2, 0xb0, 0xea, 0xa9, 0x36,
	0xe9, 0xdf, 0x5a, 0xdc, 0x12, 0x7a, 0xe7, 0xe4, 0xfb, 0xaf, 0x2f, 0xf5, 0x80, 0xec, 0xf2, 0xea,
	0x25, 0xed, 0x83, 0x01, 0xfe, 0xbe, 0x4c, 0xf9, 0x03, 0xf9, 0x88, 0xf0, 0xc6, 0x32, 0x19, 0x72,
	0xf7, 0x62, 0xd9, 0xca, 0x39, 0x36, 0xf7, 0x2e, 0x6b, 0xf3, 0x0e, 0xda, 0xd6, 0x01, 0x25, 0xad,
	0x3f, 0x1c, 0x2c, 0x63, 0xf4, 0x2e, 0x3a, 0xaf, 0xbf, 0x4d, 0x03, 0x74, 0x3a, 0x0d, 0xd0, 0xcf,
	0x69, 0x80, 0x3e, 0xcf, 0x82, 0xda, 0xe9, 0x2c, 0xa8, 0xfd, 0x98, 0x05, 0xb5, 0x37, 0x87, 0x32,
	0x36, 0x83, 0x51, 0x97, 0xf5, 0xd4, 0x90, 0x3f, 0x75, 0x2a, 0x4e, 0xec, 0xa1, 0xee, 0xbf, 0xe5,
	0x52, 0x25, 0x90, 0x4a, 0xee, 0x5f, 0xd0, 0xbb, 0xf3, 0x05, 0x66, 0x92, 0x09, 0xdd, 0xfd, 0xdf,
	0xbe, 0x8b, 0x47, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x1d, 0x3e, 0x9b, 0xa7, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Return an arbitrary vstorage datum.
	Data(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error)
	// Return the children of a given vstorage path.
	Children(ctx context.Context, in *QueryChildrenRequest, opts ...grpc.CallOption) (*QueryChildrenResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Data(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error) {
	out := new(QueryDataResponse)
	err := c.cc.Invoke(ctx, "/agoric.vstorage.Query/Data", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Children(ctx context.Context, in *QueryChildrenRequest, opts ...grpc.CallOption) (*QueryChildrenResponse, error) {
	out := new(QueryChildrenResponse)
	err := c.cc.Invoke(ctx, "/agoric.vstorage.Query/Children", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Return an arbitrary vstorage datum.
	Data(context.Context, *QueryDataRequest) (*QueryDataResponse, error)
	// Return the children of a given vstorage path.
	Children(context.Context, *QueryChildrenRequest) (*QueryChildrenResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Data(ctx context.Context, req *QueryDataRequest) (*QueryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Data not implemented")
}
func (*UnimplementedQueryServer) Children(ctx context.Context, req *QueryChildrenRequest) (*QueryChildrenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Children not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Data_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Data(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agoric.vstorage.Query/Data",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Data(ctx, req.(*QueryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Children_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChildrenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Children(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agoric.vstorage.Query/Children",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Children(ctx, req.(*QueryChildrenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agoric.vstorage.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Data",
			Handler:    _Query_Data_Handler,
		},
		{
			MethodName: "Children",
			Handler:    _Query_Children_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agoric/vstorage/query.proto",
}

func (m *QueryDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChildrenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChildrenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChildrenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChildrenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChildrenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChildrenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Children) > 0 {
		for iNdEx := len(m.Children) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Children[iNdEx])
			copy(dAtA[i:], m.Children[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Children[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryChildrenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryChildrenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, s := range m.Children {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChildrenRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChildrenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChildrenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChildrenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChildrenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChildrenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)