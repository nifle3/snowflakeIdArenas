package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Etcd struct {
}

type App struct {
	EnvType    string    `env:"ENV_TYPE"`
	StartEpoch time.Time `env:"START_EPOCH" env-layout:"2006-01-02T15:04:05Z07:00"`
	Etcd
}

func MustNew() *App {
	config := new(App)
	if err := cleanenv.ReadEnv(config); err != nil {
		panic(fmt.Sprintf("Panic during config read with error: %s", err.Error()))
	}

	return config
}
