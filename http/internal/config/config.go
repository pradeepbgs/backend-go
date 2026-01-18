package config

import "os"


type Config struct{
	Port string
	DB_URL string
}

func Load() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}
}