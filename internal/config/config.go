package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
	}

	App struct {
		Name     string `env-required:"true" yaml:"name"`
		AppURL   string `env-required:"true" yaml:"appUrl"`
		APIURL   string `env-required:"true" yaml:"apiUrl"`
		ImageURL string `env-required:"true" yaml:"imageUrl"`
		File     string `env-required:"true" yaml:"file"`
	}
)

func New(path string) *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("config new readconfig: %w", err))
	}

	return cfg
}
