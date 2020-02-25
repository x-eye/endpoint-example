package storage

import (
	"encoding/json"
	pb "example/pkg/pb/example"

	"github.com/Shopify/sarama"
)

type kafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(topic string, addrs ...string) (*kafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Metadata.Retry.Max = 1000
	p, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}
	return &kafkaProducer{producer: p, topic: topic}, nil
}

func (p *kafkaProducer) Send(request *pb.ExampleRequest) error {
	value, err := json.Marshal(request)
	if err != nil {
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(value),
	}
	_, _, err = p.producer.SendMessage(&msg)
	if err != nil {
		return err
	}
	return nil
}
