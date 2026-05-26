package core_config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string 
	PostgresUser string
	PostgresPassword string
	PostgresDB string
	RedisAddr string
	RedisPassword string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Host: os.Getenv("HOST"),
		RedisAddr: os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresDB: os.Getenv("POSTGRES_DB"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func (c *Config) Validation() error {
	cfg := map[string]string{
		"HOST": os.Getenv("HOST"),
		"REDIS_HOST": os.Getenv("REDIS_HOST"),
		"POSTGRES_USER": os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"POSTGRES_DB": os.Getenv("POSTGRES_DB"),
	}

	for _, value := range cfg {
		if value == "" {
			panic("error")
		}
	}

	return nil
}