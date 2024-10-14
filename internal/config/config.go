package config

import (
	"log"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

// Config represents the structure of the config file.
// It contains the project directory path, a list of categories, and a list of projects.
type Config struct {
	ProjectDirectory string     `yaml:"project_directory"`
	Categories       []Category `yaml:"categories"`
	Projects         []Project  `yaml:"projects"`
}

// Project represents the structure of a project.
// It contains the name, category, origin, and state of the project.
type Project struct {
	Name     string `yaml:"name"`
	ID       int    `yaml:"ID"`
	Category string `yaml:"category"`
	Origin   string `yaml:"origin"`
	State    string `yaml:"state"`
}

// Category represents the structure of a category.
// It contains the name and relative path of the category.
type Category struct {
	Name    string `yaml:"name"`
	RelPath string `yaml:"rel_path"`
}

// getConfigPath returns the path to the config file.
// exported = no
func getConfigPath() string {
	config, err := xdg.ConfigFile("pdm/config.yaml")
	if err != nil {
		log.Fatalf("Failed to get config path: %v", err)
	}
	return config
}

// WriteConfig writes the config to the config file.
// It takes a pointer to a Config struct as an argument.
// exported = yes
func WriteConfig(conf *Config) {
	data, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}
	err = os.WriteFile(getConfigPath(), data, 0o644)
	if err != nil {
		log.Fatalf("WriteFile: %v", err)
	}
}

// ReadConfig reads the config from the config file.
// It returns a pointer to a Config struct.
// exported = yes
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

func GetNewID(conf *Config) int {
	if len(conf.Projects) == 0 {
		return 1
	}
	lastID := conf.Projects[len(conf.Projects)-1].ID
	return lastID + 1
}

func AddProjectToConfig(config *Config, projectName string, projectOrigin string, projectCategory string) {
	newProject := Project{
		Name:     projectName,
		Category: projectCategory,
		Origin:   projectOrigin,
		ID:       GetNewID(config),
		State:    "new",
	}
	config.Projects = append(config.Projects, newProject)
	WriteConfig(config)
}
