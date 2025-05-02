package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type LogConfig struct {
	Level string `yaml:"level"`
}

type Config struct {
	Log      LogConfig     `yaml:"log"`
	Database StorageConfig `yaml:"db"`
	Server   ServerConfig  `yaml:"server"`
}

type ServerConfig struct {
	Address string        `yaml:"address"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func MustLoad() *Config {
	if err := loadEnv(); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	configPath := getEnv("CONFIG_PATH", "./config/local.yaml")

	serverPort := getEnv("SERVER_PORT", "8080")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	cfg.Server.Port = serverPort

	return &cfg
}

func loadEnv() error {
	var err error

	if err = godotenv.Load(); err == nil {
		return nil
	}

	return err
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
