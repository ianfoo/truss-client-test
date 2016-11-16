package clienthandler

import (
	pb "github.com/hasian/truss-client-test/printer/printer-service"
)

// Print implements Service.
func Print(NamePrint string) (*pb.PrintRequest, error) {

	request := pb.PrintRequest{
		Name: NamePrint,
	}
	return &request, nil
}
