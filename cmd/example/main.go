package main

import (
	"example/configs"
	"example/internal/services"
	"example/internal/validators"
	"log"
	"net/http"
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

	service := services.NewExampleService(validator)

	s := server.NewServer(cfg.Port)
	err = s.Run(service)
	if err != nil {
		log.Fatal(err)
	}
}
