// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppingCart.proto

package shoppingCart

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AddProductRequest struct {
	ProductID            string   `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductRequest) Reset()         { *m = AddProductRequest{} }
func (m *AddProductRequest) String() string { return proto.CompactTextString(m) }
func (*AddProductRequest) ProtoMessage()    {}
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55eb223b9cf8f13, []int{0}
}

func (m *AddProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductRequest.Unmarshal(m, b)
}
func (m *AddProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductRequest.Marshal(b, m, deterministic)
}
func (m *AddProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductRequest.Merge(m, src)
}
func (m *AddProductRequest) XXX_Size() int {
	return xxx_messageInfo_AddProductRequest.Size(m)
}
func (m *AddProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductRequest proto.InternalMessageInfo

func (m *AddProductRequest) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

type AddProductResponse struct {
	Succeed              bool     `protobuf:"varint,1,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductResponse) Reset()         { *m = AddProductResponse{} }
func (m *AddProductResponse) String() string { return proto.CompactTextString(m) }
func (*AddProductResponse) ProtoMessage()    {}
func (*AddProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55eb223b9cf8f13, []int{1}
}

func (m *AddProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductResponse.Unmarshal(m, b)
}
func (m *AddProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductResponse.Marshal(b, m, deterministic)
}
func (m *AddProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductResponse.Merge(m, src)
}
func (m *AddProductResponse) XXX_Size() int {
	return xxx_messageInfo_AddProductResponse.Size(m)
}
func (m *AddProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductResponse proto.InternalMessageInfo

func (m *AddProductResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

type GetShoppingCartResponse struct {
	ProductID            []string `protobuf:"bytes,1,rep,name=ProductID,proto3" json:"ProductID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingCartResponse) Reset()         { *m = GetShoppingCartResponse{} }
func (m *GetShoppingCartResponse) String() string { return proto.CompactTextString(m) }
func (*GetShoppingCartResponse) ProtoMessage()    {}
func (*GetShoppingCartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55eb223b9cf8f13, []int{2}
}

func (m *GetShoppingCartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingCartResponse.Unmarshal(m, b)
}
func (m *GetShoppingCartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingCartResponse.Marshal(b, m, deterministic)
}
func (m *GetShoppingCartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingCartResponse.Merge(m, src)
}
func (m *GetShoppingCartResponse) XXX_Size() int {
	return xxx_messageInfo_GetShoppingCartResponse.Size(m)
}
func (m *GetShoppingCartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingCartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingCartResponse proto.InternalMessageInfo

func (m *GetShoppingCartResponse) GetProductID() []string {
	if m != nil {
		return m.ProductID
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55eb223b9cf8f13, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddProductRequest)(nil), "shoppingCart.AddProductRequest")
	proto.RegisterType((*AddProductResponse)(nil), "shoppingCart.AddProductResponse")
	proto.RegisterType((*GetShoppingCartResponse)(nil), "shoppingCart.GetShoppingCartResponse")
	proto.RegisterType((*Empty)(nil), "shoppingCart.Empty")
}

func init() { proto.RegisterFile("shoppingCart.proto", fileDescriptor_e55eb223b9cf8f13) }

var fileDescriptor_e55eb223b9cf8f13 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xce, 0xc8, 0x2f,
	0x28, 0xc8, 0xcc, 0x4b, 0x77, 0x4e, 0x2c, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x32, 0xe4, 0x12, 0x74, 0x4c, 0x49, 0x09, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e,
	0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe1, 0xe2, 0x84, 0x8a, 0x78, 0xba, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x94, 0xf4, 0xb8, 0x84, 0x90, 0xb5, 0x14, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x07, 0x97, 0x26, 0x27, 0xa7, 0xa6, 0xa6, 0x80,
	0x75, 0x70, 0x04, 0xc1, 0xb8, 0x4a, 0xe6, 0x5c, 0xe2, 0xee, 0xa9, 0x25, 0xc1, 0x48, 0xb6, 0xc2,
	0x35, 0xa1, 0x59, 0xc4, 0x8c, 0x6a, 0x11, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0xd1,
	0x56, 0x46, 0x2e, 0x61, 0x64, 0xfd, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xfe, 0x5c,
	0x5c, 0x08, 0x97, 0x08, 0xc9, 0xeb, 0xa1, 0xf8, 0x16, 0xc3, 0x5b, 0x52, 0x0a, 0xb8, 0x15, 0x40,
	0xdd, 0xe3, 0xcb, 0xc5, 0x8f, 0xe6, 0x54, 0x21, 0x61, 0x54, 0x4d, 0x60, 0x07, 0x49, 0xa9, 0xa2,
	0x0a, 0xe2, 0xf0, 0x5e, 0x12, 0x1b, 0x38, 0xc4, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x9b, 0x43, 0xb0, 0x87, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShoppingCartServiceClient is the client API for ShoppingCartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingCartServiceClient interface {
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	GetShoppingCart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetShoppingCartResponse, error)
}

type shoppingCartServiceClient struct {
	cc *grpc.ClientConn
}

func NewShoppingCartServiceClient(cc *grpc.ClientConn) ShoppingCartServiceClient {
	return &shoppingCartServiceClient{cc}
}

func (c *shoppingCartServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, "/shoppingCart.ShoppingCartService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingCartServiceClient) GetShoppingCart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetShoppingCartResponse, error) {
	out := new(GetShoppingCartResponse)
	err := c.cc.Invoke(ctx, "/shoppingCart.ShoppingCartService/GetShoppingCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingCartServiceServer is the server API for ShoppingCartService service.
type ShoppingCartServiceServer interface {
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	GetShoppingCart(context.Context, *Empty) (*GetShoppingCartResponse, error)
}

// UnimplementedShoppingCartServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingCartServiceServer struct {
}

func (*UnimplementedShoppingCartServiceServer) AddProduct(ctx context.Context, req *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedShoppingCartServiceServer) GetShoppingCart(ctx context.Context, req *Empty) (*GetShoppingCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingCart not implemented")
}

func RegisterShoppingCartServiceServer(s *grpc.Server, srv ShoppingCartServiceServer) {
	s.RegisterService(&_ShoppingCartService_serviceDesc, srv)
}

func _ShoppingCartService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppingCart.ShoppingCartService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingCartService_GetShoppingCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingCartServiceServer).GetShoppingCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppingCart.ShoppingCartService/GetShoppingCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingCartServiceServer).GetShoppingCart(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingCartService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shoppingCart.ShoppingCartService",
	HandlerType: (*ShoppingCartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ShoppingCartService_AddProduct_Handler,
		},
		{
			MethodName: "GetShoppingCart",
			Handler:    _ShoppingCartService_GetShoppingCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppingCart.proto",
}