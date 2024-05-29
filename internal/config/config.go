package config

import (
	"log"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ProjectDirectory string `yaml:"project_directory"`
}

func getConfigPath() string {
	config, err := xdg.ConfigFile("pdm/config.yaml")
	if err != nil {
		log.Fatalf("Failed to get config path: %v", err)
	}
	return config
}

func ReadConfig() *Config {
	yamlFile, err := os.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &conf
}
