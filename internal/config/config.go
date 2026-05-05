package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port string `env:"PORT" env-default:"8080"`
	Env  string `env:"ENV"  env-default:"development"`
}

func MustRead() *Config {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		panic(err)
	}
	return cfg
}
