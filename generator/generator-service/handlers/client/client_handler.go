package clienthandler

import (
	pb "github.com/hasian/truss-client-test/generator/generator-service"
)

// Generate implements Service.
func Generate(NameGenerate string) (*pb.GenerateRequest, error) {

	request := pb.GenerateRequest{
		Name: NameGenerate,
	}
	return &request, nil
}
