package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Etcd struct {
}

type Grpc struct {
}

type App struct {
	EnvType string `env:"ENV-TYPE"`
	Etcd
	Grpc
}

func MustNew() *App {
	config := new(App)
	if err := cleanenv.ReadEnv(config); err != nil {
		panic(fmt.Sprintf("Panic during config read with error: %s", err.Error()))
	}

	return config
}
