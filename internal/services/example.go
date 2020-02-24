package services

import (
	"context"
	pb "example/pkg/pb/example"

	"github.com/utrack/clay/v2/transport"
)

type exampleService struct {
}

func (s *exampleService) GetDescription() transport.ServiceDesc {
	return pb.NewExampleServiceServiceDesc(s)
}

func NewExampleService() *exampleService {
	return &exampleService{}
}

func (s *exampleService) ExampleEndpoint(context.Context, *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	resp := pb.ExampleResponse{}
	return &resp, nil
}
