package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/hasian/truss-client-test/printer/printer-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AddServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints /*, tracer stdopentracing.Tracer, logger log.Logger*/) pb.PrinterServer {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerErrorLogger(logger),
	//}
	return &grpcServer{
		// printer

		print: grpctransport.NewServer(
			ctx,
			endpoints.PrintEndpoint,
			DecodeGRPCPrintRequest,
			EncodeGRPCPrintResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "Print", logger)))...,
		),
	}
}

type grpcServer struct {
	print grpctransport.Handler
}

// Methods

func (s *grpcServer) Print(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	_, rep, err := s.print.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PrintResponse), nil
}

// Server Decode

// DecodeGRPCPrintRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC print request to a user-domain print request. Primarily useful in a server.
func DecodeGRPCPrintRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.PrintRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCPrintResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC print reply to a user-domain print response. Primarily useful in a client.
func DecodeGRPCPrintResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.PrintResponse)
	return reply, nil
}

// Server Encode

// EncodeGRPCPrintResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain print response to a gRPC print reply. Primarily useful in a server.
func EncodeGRPCPrintResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.PrintResponse)
	return resp, nil
}

// Client Encode

// EncodeGRPCPrintRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain print request to a gRPC print request. Primarily useful in a client.
func EncodeGRPCPrintRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PrintRequest)
	return req, nil
}
