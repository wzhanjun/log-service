package client

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var Cfg Env

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	if err := env.Parse(&Cfg); err != nil {
		log.Println(err)
	}

	log.Printf("%#v\n", Cfg)
}

type Env struct {
	AppId             string `env:"APP_ID"`
	AppDeBug          bool   `env:"APP_DEBUG"`
	LogServiceAddress string `env:"LOGSERVICE_ADDRESS"`
	LogServiceEsIndex string `env:"LOGSERVICE_ESINDEX"`
}
