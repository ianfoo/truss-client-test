// Package grpc provides a gRPC client for the add service.
package grpc

import (
	//"time"

	//jujuratelimit "github.com/juju/ratelimit"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"

	//"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/ratelimit"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/hasian/truss-client-test/generator/generator-service"
	svc "github.com/hasian/truss-client-test/generator/generator-service/generated"
	handler "github.com/hasian/truss-client-test/generator/generator-service/handlers/server"
)

// New returns an AddService backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn /*, tracer stdopentracing.Tracer, logger log.Logger*/) handler.Service {
	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.

	//limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))

	var generateEndpoint endpoint.Endpoint
	{
		generateEndpoint = grpctransport.NewClient(
			conn,
			"generator.Generator",
			"Generate",
			svc.EncodeGRPCGenerateRequest,
			svc.DecodeGRPCGenerateResponse,
			pb.GenerateResponse{},
			//grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "Generate", logger)),
		).Endpoint()
		//generateEndpoint = opentracing.TraceClient(tracer, "Generate")(generateEndpoint)
		//generateEndpoint = limiter(generateEndpoint)
		//generateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//Name:    "Generate",
		//Timeout: 30 * time.Second,
		//}))(generateEndpoint)
	}

	return svc.Endpoints{

		GenerateEndpoint: generateEndpoint,
	}
}
