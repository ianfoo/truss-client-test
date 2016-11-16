package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"
    "fmt"
    "log"
    "os"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	pb "github.com/hasian/truss-client-test/printer/printer-service"
    genclient "github.com/hasian/truss-client-test/generator/generator-service/generated/client/http"
    gensvc "github.com/hasian/truss-client-test/generator/generator-service/handlers/server"
    genpb  "github.com/hasian/truss-client-test/generator/generator-service"
)

var gen gensvc.Service

// NewService returns a naïve, stateless implementation of Service.
func NewService() Service {
	return printerService{}
}

type printerService struct{}

// Print implements Service.
func (s printerService) Print(ctx context.Context, in *pb.PrintRequest) (*pb.PrintResponse, error) {
	_ = ctx
	_ = in
    name := in.Name
    if name == "" {
        name = "SOMEBODY:Printer"
    }
    gresp, err := gen.Generate(context.Background(), &genpb.GenerateRequest{
        Name: name,
    })
    if err != nil {
        return &pb.PrintResponse{
            Message: fmt.Sprintf("error generating message: %v", err),
        }, nil
    }
	response := pb.PrintResponse{
	    Message: gresp.Greeting + "!",
	    GeneratedAt: gresp.GeneratedAt,
	}
	return &response, nil
}

type Service interface {
	Print(ctx context.Context, in *pb.PrintRequest) (*pb.PrintResponse, error)
}

func init() {
    log.Println("starting printer service")
    genaddr := os.Getenv("GENERATOR_ADDR")
    if genaddr == "" {
        genaddr = "localhost:10000"
        log.Printf("Trying %s for generator address", genaddr)
    }
    var err error
    gen, err = genclient.New(genaddr)
    if err != nil {
        log.Fatalf("error creating generator client: %v", err)
    }
}
