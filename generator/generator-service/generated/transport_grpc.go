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
	pb "github.com/hasian/truss-client-test/generator/generator-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AddServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints /*, tracer stdopentracing.Tracer, logger log.Logger*/) pb.GeneratorServer {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerErrorLogger(logger),
	//}
	return &grpcServer{
		// generator

		generate: grpctransport.NewServer(
			ctx,
			endpoints.GenerateEndpoint,
			DecodeGRPCGenerateRequest,
			EncodeGRPCGenerateResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "Generate", logger)))...,
		),
	}
}

type grpcServer struct {
	generate grpctransport.Handler
}

// Methods

func (s *grpcServer) Generate(ctx context.Context, req *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	_, rep, err := s.generate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenerateResponse), nil
}

// Server Decode

// DecodeGRPCGenerateRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC generate request to a user-domain generate request. Primarily useful in a server.
func DecodeGRPCGenerateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GenerateRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCGenerateResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC generate reply to a user-domain generate response. Primarily useful in a client.
func DecodeGRPCGenerateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GenerateResponse)
	return reply, nil
}

// Server Encode

// EncodeGRPCGenerateResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain generate response to a gRPC generate reply. Primarily useful in a server.
func EncodeGRPCGenerateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GenerateResponse)
	return resp, nil
}

// Client Encode

// EncodeGRPCGenerateRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain generate request to a gRPC generate request. Primarily useful in a client.
func EncodeGRPCGenerateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GenerateRequest)
	return req, nil
}
