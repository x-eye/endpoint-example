package services

import (
	"context"
	pb "example/pkg/pb/example"

	"github.com/utrack/clay/v2/transport"
)

type IValidator interface {
	Validate(*pb.ExampleRequest) error
}

type ISender interface {
	Send(*pb.ExampleRequest) error
}

type exampleService struct {
	validator IValidator
	sender    ISender
}

func NewExampleService(validator IValidator, sender ISender) *exampleService {
	return &exampleService{validator: validator, sender: sender}
}

func (s *exampleService) GetDescription() transport.ServiceDesc {
	return pb.NewExampleServiceServiceDesc(s)
}

func (s *exampleService) ExampleEndpoint(_ context.Context, request *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	err := s.validator.Validate(request)
	if err != nil {
		return nil, err
	}
	err = s.sender.Send(request)
	if err != nil {
		return nil, err
	}
	resp := pb.ExampleResponse{}
	return &resp, nil
}
