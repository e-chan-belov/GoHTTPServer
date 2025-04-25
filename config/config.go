package config

import (
	"flag"
	"github.com/numbergroup/cleanenv"
	"log"
	"os"
)

type HTTPConfig struct {
	Address string
}

func MustLoad(cfgPath string, cfg any) {
	if cfgPath == "" {
		log.Fatal("Config path is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist by this path: %s", cfgPath)
	}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}
}

type AppFlags struct {
	ConfigPath string
}

func ParseFlags() AppFlags {
	configPath := flag.String("config", "", "Path to config")
	flag.Parse()
	return AppFlags{
		ConfigPath: *configPath,
	}
}
