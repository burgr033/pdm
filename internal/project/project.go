package project

import (
	"github.com/burgr033/pdm/internal/config"
	gitLib "github.com/burgr033/pdm/internal/project/git"
)

// GetProjectTable returns a table of all projects.
// It returns the project name, category, origin, and state.
// exported = yes
func GetProjectTable() [][]string {
	conf := config.ReadConfig()
	var projects [][]string
	for _, p := range conf.Projects {
		projects = append(projects, []string{p.Name, p.Category, p.Origin, p.State})
	}
	return projects
}

// AddProject adds a project to the configuration.
// exported = yes
func AddProject(name string, category string, origin string, state string) {
	conf := config.ReadConfig()
	conf.Projects = append(conf.Projects, config.Project{Name: name, Category: category, Origin: origin, State: state})
	config.WriteConfig(conf)
}

// CloneProject clones a project from a URL.
// It clones the project and adds it to the configuration.
// exported = yes
func CloneProject(url string, category string) {
	projectName := gitLib.GetRepoNameFromURL(url)
	gitLib.CloneAction(url, category)
	AddProject(projectName, category, url, "")
}

// InitProject initializes a project in the current directory.
// It initializes a git repository and adds it to the configuration.
// exprted = yes
func InitProject(name string, category string) {
	gitLib.InitAction(name, category)
	AddProject(name, category, "local", "")
}
