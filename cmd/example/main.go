package main

import (
	"example/configs"
	"example/internal/services"
	"log"

	"github.com/utrack/clay/v2/transport/server"
)

func main() {
	cfg := configs.Cfg()
	s := server.NewServer(cfg.Port)
	service := services.NewExampleService()
	err := s.Run(service)
	if err != nil {
		log.Fatal(err)
	}
}
