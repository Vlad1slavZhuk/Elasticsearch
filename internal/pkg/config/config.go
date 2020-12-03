package config

import (
	"log"
	"os"
)

type Config struct {
	PortHTTP string
}

var port string

func init() {
	if port = os.Getenv("SERVER_PORT"); port == "" {
		log.Fatal("Set SERVER_PORT!")
	}
}

func NewConfigFromEnv() *Config {
	config := new(Config)
	config.PortHTTP = ":" + port
	return config
}
