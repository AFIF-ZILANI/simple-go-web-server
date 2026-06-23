package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Host    string `yaml:"host" env-required:"true"`
	Port    int    `yaml:"port" env-required:"true"`
	Address string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}

func MustLoadConfig() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("CONFIG_PATH environment variable or --config flag must be set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %s", err.Error())
	}

	return &cfg
}
