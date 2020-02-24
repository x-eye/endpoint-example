package services

import (
	"context"
	pb "example/pkg/pb/example"

	"github.com/utrack/clay/v2/transport"
)

type IValidator interface {
	Validate(*pb.ExampleRequest) error
}

type exampleService struct {
}

func (s *exampleService) GetDescription() transport.ServiceDesc {
	return pb.NewExampleServiceServiceDesc(s)
}

func NewExampleService() *exampleService {
	return &exampleService{}
}

func (s *exampleService) ExampleEndpoint(_ context.Context, request *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	if err := s.validator.Validate(request); err != nil {
		return nil, err
	}
	resp := pb.ExampleResponse{}
	return &resp, nil
}
