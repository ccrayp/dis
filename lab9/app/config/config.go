package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	PostgreSQL PostgreSQL `yaml:"db"`
}

type PostgreSQL struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SslMode  string `yaml:"sslmode"`
}

func LoadConfig() (*Config, error) {
	var config Config

	if err := cleanenv.ReadConfig("config.yaml", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
