package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port int `envconfig:"PORT" default:"9003"`
}

var cfg Config

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
}

func Cfg() Config {
	return cfg
}
