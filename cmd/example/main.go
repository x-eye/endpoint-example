package main

import (
	"context"
	"example/configs"
	"example/internal/services"
	"example/internal/storage"
	"example/internal/validators"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"

	"github.com/utrack/clay/v2/transport/server"
)

func main() {
	cfg := configs.Cfg()
	id_re, err := regexp.Compile(cfg.IdRegexp)
	if err != nil {
		log.Fatal(err)
	}
	validator := validators.NewExampleValidator(id_re)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill, os.Interrupt)
	go func() {
		<-done
		cancel()
	}()

	partner := storage.NewPartner(http.DefaultClient, cfg.PartnerUrl, cfg.PartnerQueueMaxLength)
	go partner.Consume(ctx)
	producer, err := storage.NewKafkaProducer(cfg.KafkaTopic, cfg.KafkaAddr)
	if err != nil {
		log.Fatal(err)
	}
	senders := storage.NewSenderFanout(partner, producer)
	service := services.NewExampleService(validator, senders)

	s := server.NewServer(cfg.Port)
	err = s.Run(service)
	if err != nil {
		log.Fatal(err)
	}
}
