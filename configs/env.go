package configs

import "github.com/kelseyhightower/envconfig"

const ENV_PREFIX = ""

type Config struct {
	Port       int    `envconfig:"PORT" default:"9003"`
	IdRegexp   string `envconfig:"ID_REGEXP" default:"^([0-9A-Fa-f]{2}[-]){5}([0-9A-Fa-f]{2})$"`
	PartnerUrl string `envconfig:"PARTNER_URL"`

	PartnerQueueMaxLength int `envconfig:"PARTNER_QUEUE_MAX_LENGTH" default:"1000"`
}

var cfg Config

func init() {
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		panic(err)
	}
}

func Cfg() Config {
	return cfg
}
