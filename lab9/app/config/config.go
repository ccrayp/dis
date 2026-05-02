package config

import "github.com/ilyakaznacheev/cleanenv"

// Структура конфигурации приложение
type Config struct {
	PostgreSQL PostgreSQL `yaml:"db"`
}

// Структура конфиграции параметров для подключения к базе данных
type PostgreSQL struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SslMode  string `yaml:"sslmode"`
}

// Функция загрузки конфигурации приложения
func LoadConfig() (*Config, error) {
	var config Config

	if err := cleanenv.ReadConfig("config.yaml", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
