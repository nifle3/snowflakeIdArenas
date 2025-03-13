package server

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}

func mustNewConfig() *config {
	config := new(config)
	if err := cleanenv.ReadEnv(config); err != nil {
		slog.Error("Server config read err", slog.String("Err", err.Error()))
		panic("Server config read err")
	}

	return config
}
