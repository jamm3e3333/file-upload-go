package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Port uint32 `env:"HTTP_LISTEN_PORT" env-default:"8889"`
}

func NewAppCfg() *App {
	cfg := &App{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("failed to create app config: %s", err.Error())
	}
	return cfg
}
