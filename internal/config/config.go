package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseName string `default:"tasks.db" split_words:"true"`
}

func Init() *Config {
	var c Config

	err := envconfig.Process("task", &c)
	if err != nil {
		log.Fatal(err)
	}
	return &c
}
