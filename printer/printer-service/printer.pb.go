// Code generated by protoc-gen-go.
// source: printer.proto
// DO NOT EDIT!

/*
Package printer is a generated protocol buffer package.

It is generated from these files:
	printer.proto

It has these top-level messages:
	PrintRequest
	PrintResponse
*/
package printer

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

type PrintRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PrintRequest) Reset()                    { *m = PrintRequest{} }
func (m *PrintRequest) String() string            { return proto.CompactTextString(m) }
func (*PrintRequest) ProtoMessage()               {}
func (*PrintRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PrintResponse struct {
	Message     string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	GeneratedAt string `protobuf:"bytes,2,opt,name=generated_at,json=generatedAt" json:"generated_at,omitempty"`
}

func (m *PrintResponse) Reset()                    { *m = PrintResponse{} }
func (m *PrintResponse) String() string            { return proto.CompactTextString(m) }
func (*PrintResponse) ProtoMessage()               {}
func (*PrintResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*PrintRequest)(nil), "printer.PrintRequest")
	proto.RegisterType((*PrintResponse)(nil), "printer.PrintResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Printer service

type PrinterClient interface {
	Print(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintResponse, error)
}

type printerClient struct {
	cc *grpc.ClientConn
}

func NewPrinterClient(cc *grpc.ClientConn) PrinterClient {
	return &printerClient{cc}
}

func (c *printerClient) Print(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintResponse, error) {
	out := new(PrintResponse)
	err := grpc.Invoke(ctx, "/printer.Printer/Print", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Printer service

type PrinterServer interface {
	Print(context.Context, *PrintRequest) (*PrintResponse, error)
}

func RegisterPrinterServer(s *grpc.Server, srv PrinterServer) {
	s.RegisterService(&_Printer_serviceDesc, srv)
}

func _Printer_Print_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServer).Print(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/printer.Printer/Print",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServer).Print(ctx, req.(*PrintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Printer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "printer.Printer",
	HandlerType: (*PrinterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Print",
			Handler:    _Printer_Print_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("printer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0xcc,
	0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x64, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xca, 0x94, 0x94, 0xb8, 0x78, 0x02, 0x40, 0x0a, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x1f, 0x2e, 0x5e, 0xa8, 0x9a, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x3a, 0x18, 0x57,
	0x48, 0x91, 0x8b, 0x27, 0x3d, 0x35, 0x2f, 0xb5, 0x28, 0xb1, 0x24, 0x35, 0x25, 0x3e, 0xb1, 0x44,
	0x82, 0x09, 0x2c, 0xcd, 0x0d, 0x17, 0x73, 0x2c, 0x31, 0x0a, 0xe4, 0x62, 0x0f, 0x80, 0x38, 0x4d,
	0xc8, 0x8d, 0x8b, 0x15, 0xcc, 0x14, 0x12, 0xd5, 0x83, 0x39, 0x1e, 0xd9, 0x31, 0x52, 0x62, 0xe8,
	0xc2, 0x10, 0xfb, 0x95, 0xf8, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0x21, 0xc4, 0xa6, 0x0f, 0x16,
	0x4f, 0x62, 0x03, 0xfb, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x0f, 0xf3, 0xb4, 0x03,
	0x01, 0x00, 0x00,
}
