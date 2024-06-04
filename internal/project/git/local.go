package git

import (
	"log"
	"path"

	"github.com/burgr033/pdm/internal/config"
	"github.com/go-git/go-git/v5"
)

// InitGitRepository initializes a git repository in the current directory.
// exported = yes
func InitAction(name string, category string) {
	conf := config.ReadConfig()
	repositoryPath := path.Join(conf.ProjectDirectory, category, name)
	_, err := git.PlainInit(repositoryPath, false)
	if err != nil {
		log.Fatalf("Failed to initialize git repository: %v", err)
	}
}
