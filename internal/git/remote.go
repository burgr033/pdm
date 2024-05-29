package git

import (
	"log"
	"net/url"
	"os"

	"github.com/burgr033/pdm/internal/config"
	"github.com/go-git/go-git/v5"
)

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getRepoNameFromURL(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %s", URL)
	}
	return u.Path
}

func CloneAction(URL string) {
	if !IsUrl(URL) {
		log.Fatalf("Invalid URL: %s", URL)
	}
	conf := config.ReadConfig()
	repositoryPath := conf.ProjectDirectory + getRepoNameFromURL(URL)
	_, err := git.PlainClone(repositoryPath, false, &git.CloneOptions{
		URL:      URL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatalf("Failed to clone repository: %s to %s | %v", URL, repositoryPath, err)
	}
}
