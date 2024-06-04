package git

import (
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/burgr033/pdm/internal/config"
	"github.com/go-git/go-git/v5"
)

// IsUrl checks if a string is a valid URL.
// exported = no
func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// getRepoNameFromURL returns the repository name from a URL.
// exported = no
func GetRepoNameFromURL(URL string) string {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %s", URL)
	}
	path := u.Path
	// omit the .git if exist
	if strings.HasSuffix(path, ".git") {
		path = path[:len(path)-4]
	}
	// get the last part of the path
	path = path[strings.LastIndex(path, "/")+1:]

	return path
}

// CloneAction clones the repository from the given URL
// exported = yes
func CloneAction(url string, category string) {
	if !isUrl(url) {
		log.Fatalf("Invalid URL: %s", url)
	}
	conf := config.ReadConfig()
	repositoryPath := path.Join(conf.ProjectDirectory, category, GetRepoNameFromURL(url))
	_, err := git.PlainClone(repositoryPath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatalf("Failed to clone repository: %s to %s | %v", url, repositoryPath, err)
	}
}
