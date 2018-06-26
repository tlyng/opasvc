// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/echo.proto

package pb // import "github.com/tlyng/opasvc/pb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Greeting             string   `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_c94f05759dfb829e, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_c94f05759dfb829e, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "poc.opasvc.Request")
	proto.RegisterType((*Response)(nil), "poc.opasvc.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/poc.opasvc.Hello/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Say(context.Context, *Request) (*Response, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poc.opasvc.Hello/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Say(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poc.opasvc.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/echo.proto",
}

func init() { proto.RegisterFile("api/echo.proto", fileDescriptor_echo_c94f05759dfb829e) }

var fileDescriptor_echo_c94f05759dfb829e = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0xcf, 0x82, 0x30,
	0x10, 0x86, 0xc3, 0xf7, 0xa9, 0xe0, 0x0d, 0x0e, 0xd5, 0x81, 0x10, 0x07, 0x43, 0x1c, 0x9c, 0x5a,
	0xa3, 0x13, 0xab, 0x93, 0x33, 0x6e, 0x6e, 0xa5, 0xb9, 0x14, 0x12, 0xe8, 0x55, 0x5a, 0x4c, 0xf8,
	0xf7, 0x26, 0x20, 0x1a, 0xb7, 0x7b, 0x9f, 0xdc, 0xdd, 0x93, 0x17, 0x56, 0xd2, 0x56, 0x02, 0x55,
	0x49, 0xdc, 0xb6, 0xe4, 0x89, 0x81, 0x25, 0xc5, 0xc9, 0x4a, 0xf7, 0x54, 0x69, 0x06, 0x61, 0x8e,
	0x8f, 0x0e, 0x9d, 0x67, 0x0c, 0x66, 0x46, 0x36, 0x18, 0x07, 0xbb, 0xe0, 0xb0, 0xcc, 0x87, 0x99,
	0x25, 0x10, 0xe9, 0x16, 0xd1, 0x57, 0x46, 0xc7, 0x7f, 0x03, 0xff, 0xe4, 0x74, 0x0f, 0x51, 0x8e,
	0xce, 0x92, 0x71, 0xc8, 0x62, 0x08, 0x1b, 0x74, 0x4e, 0xea, 0xe9, 0x7c, 0x8a, 0xa7, 0x0c, 0xe6,
	0x57, 0xac, 0x6b, 0x62, 0x47, 0xf8, 0xbf, 0xc9, 0x9e, 0xad, 0xf9, 0xd7, 0xce, 0xdf, 0xea, 0x64,
	0xf3, 0x0b, 0xc7, 0xa7, 0x97, 0xed, 0x3d, 0xd1, 0x95, 0x2f, 0xbb, 0x82, 0x2b, 0x6a, 0x84, 0xaf,
	0x7b, 0xa3, 0xc5, 0xb8, 0x23, 0x6c, 0x51, 0x2c, 0x86, 0x32, 0xe7, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x23, 0x5e, 0x4b, 0xde, 0x00, 0x00, 0x00,
}
