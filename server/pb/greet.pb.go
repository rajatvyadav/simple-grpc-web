// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package pb

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

type UserFullName struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFullName) Reset()         { *m = UserFullName{} }
func (m *UserFullName) String() string { return proto.CompactTextString(m) }
func (*UserFullName) ProtoMessage()    {}
func (*UserFullName) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *UserFullName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFullName.Unmarshal(m, b)
}
func (m *UserFullName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFullName.Marshal(b, m, deterministic)
}
func (m *UserFullName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFullName.Merge(m, src)
}
func (m *UserFullName) XXX_Size() int {
	return xxx_messageInfo_UserFullName.Size(m)
}
func (m *UserFullName) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFullName.DiscardUnknown(m)
}

var xxx_messageInfo_UserFullName proto.InternalMessageInfo

func (m *UserFullName) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserFullName) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type UserRequest struct {
	User                 *UserFullName `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUser() *UserFullName {
	if m != nil {
		return m.User
	}
	return nil
}

type UserResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserFullName)(nil), "pb.UserFullName")
	proto.RegisterType((*UserRequest)(nil), "pb.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe0, 0xe2, 0x09,
	0x2d, 0x4e, 0x2d, 0x72, 0x2b, 0xcd, 0xc9, 0xf1, 0x4b, 0xcc, 0x4d, 0x15, 0x92, 0xe1, 0xe2, 0x4c,
	0xcb, 0x2c, 0x2a, 0x2e, 0x01, 0x71, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10, 0x02, 0x42,
	0x52, 0x5c, 0x1c, 0x39, 0x89, 0x50, 0x49, 0x26, 0xb0, 0x24, 0x9c, 0xaf, 0x64, 0xcc, 0xc5, 0x0d,
	0x32, 0x29, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x85, 0x8b, 0xa5, 0xb4, 0x38, 0xb5,
	0x08, 0x6c, 0x06, 0xb7, 0x91, 0x80, 0x5e, 0x41, 0x92, 0x1e, 0xb2, 0x45, 0x41, 0x60, 0x59, 0x25,
	0x0d, 0x88, 0xf5, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30, 0xcb, 0x61, 0x5c, 0x23, 0x7b, 0x2e, 0x1e, 0x77, 0x90, 0xdb,
	0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xf4, 0xb9, 0x38, 0x8a, 0x13, 0x2b, 0x3d, 0x52,
	0x73, 0x72, 0xf2, 0x85, 0xf8, 0x61, 0xa6, 0x43, 0x2d, 0x97, 0x12, 0x40, 0x08, 0x40, 0x0c, 0x56,
	0x62, 0x48, 0x62, 0x03, 0x7b, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x96, 0x9f, 0x10,
	0x03, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.GreetService/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	SayHello(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) SayHello(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GreetService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}