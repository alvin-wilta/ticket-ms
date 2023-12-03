package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Env            string `env:"ENV,required"`
	ServiceAddr    string `env:"SERVICE_ADDR"`
	ServicePort    string `env:"SERVICE_PORT,required"`
	NsqAddr        string `env:"NSQLOOKUPD_ADDR"`
	NsqPort        string `env:"NSQLOOKUPD_PORT,required"`
	NsqMaxAttempts int    `env:"NSQ_MAX_ATTEMPTS"`
	NsqMaxInFlight int    `env:"NSQ_MAX_INFLIGHT"`
	PostgresAddr   string `env:"POSTGRES_ADDR"`
	PostgresPort   string `env:"POSTGRES_PORT,required"`
	PostgresUser   string `env:"POSTGRES_USER,required"`
	PostgresPass   string `env:"POSTGRES_PASS,required"`
	RedisAddr      string `env:"REDIS_ADDR"`
	RedisPort      string `env:"REDIS_PORT"`
}

func New() *Config {
	var err error
	cfg := &Config{}
	modePtr := flag.String("mode", "prod", "Set service enviroment mode, prod or dev. Default is prod.")
	flag.Parse()
	if *modePtr == "dev" {
		err = godotenv.Load(".env.dev")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		log.Fatalf("[Config] Unable to initialize environment variables: %v", err)
	}
	err = env.Parse(cfg)
	log.Print(cfg.Env)
	log.Print(cfg.NsqAddr)
	if err != nil {
		log.Fatalf("[Config] Unable to initialize env to struct: %v", err)
	}

	return cfg
}
