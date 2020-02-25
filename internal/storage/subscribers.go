package storage

import (
	"example/internal/services"
	pb "example/pkg/pb/example"
)

type senderFanout struct {
	subscribers []services.ISender
}

func NewSenderFanout(subscribers ...services.ISender) *senderFanout {
	return &senderFanout{subscribers: subscribers}
}

func (s *senderFanout) Send(request *pb.ExampleRequest) error {
	for _, sub := range s.subscribers {
		if err := sub.Send(request); err != nil {
			return err
		}
	}
	return nil
}
