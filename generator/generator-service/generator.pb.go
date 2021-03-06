// Code generated by protoc-gen-go.
// source: generator.proto
// DO NOT EDIT!

/*
Package generator is a generated protocol buffer package.

It is generated from these files:
	generator.proto

It has these top-level messages:
	GenerateRequest
	GenerateResponse
*/
package generator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hasian/truss-client-test/generator/generator-service/third_party/googleapis/google/api"

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

type GenerateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GenerateRequest) Reset()                    { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()               {}
func (*GenerateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GenerateResponse struct {
	Greeting    string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
	GeneratedAt int64  `protobuf:"varint,2,opt,name=generated_at,json=generatedAt" json:"generated_at,omitempty"`
}

func (m *GenerateResponse) Reset()                    { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()               {}
func (*GenerateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "generator.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "generator.GenerateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Generator service

type GeneratorClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := grpc.Invoke(ctx, "/generator.Generator/Generate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generator.Generator/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generator.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Generator_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("generator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0xcd, 0x4b,
	0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5,
	0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x14, 0x2a, 0xa9, 0x72, 0xf1, 0xbb, 0x43, 0x94,
	0xa6, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x81, 0x5c, 0x02, 0x08, 0x65, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0xe9, 0x45, 0xa9, 0xa9, 0x25, 0x99, 0x79,
	0xe9, 0x50, 0xb5, 0x70, 0xbe, 0x90, 0x22, 0x17, 0x0f, 0xd4, 0x05, 0xa9, 0x29, 0xf1, 0x89, 0x25,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0xdc, 0x70, 0x31, 0xc7, 0x12, 0xa3, 0x64, 0x2e, 0x4e,
	0x77, 0x98, 0x23, 0x85, 0xc2, 0xb8, 0x38, 0x60, 0xe6, 0x0b, 0x49, 0xe9, 0x21, 0x7c, 0x83, 0xe6,
	0x36, 0x29, 0x69, 0xac, 0x72, 0x10, 0x07, 0x29, 0x09, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x5b,
	0x88, 0x53, 0x1f, 0x26, 0x95, 0xc4, 0x06, 0xf6, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x04, 0x5b, 0xc9, 0x21, 0x01, 0x00, 0x00,
}
