package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppName     string `env:"APP_NAME" envDefault:"MyApp"`
	Port        int    `env:"APP_PORT" envDefault:"8080"`
	DatabaseUrl string `env:"DB_URL"`
}

var Cfg Config

func Init() {

	// Parse config
	if err := env.Parse(&Cfg); err != nil {
		log.Fatal("Failed to parse main config:", err)
	}

}
