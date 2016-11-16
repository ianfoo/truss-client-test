package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	"time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/hasian/truss-client-test/generator/generator-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() Service {
	return generatorService{}
}

type generatorService struct{}

// Generate implements Service.
func (s generatorService) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	_ = ctx
	_ = in
    name := in.Name
    if name == "" {
        name = "Somebody:Generated"
    }
	response := pb.GenerateResponse{
        Greeting: "Hello, " + name,
        GeneratedAt: time.Now().Unix(),
	}
	return &response, nil
}

type Service interface {
	Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error)
}
