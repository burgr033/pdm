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

func CloneAction(URL string) {
	if !IsUrl(URL) {
		log.Fatalf("Invalid URL: %s", URL)
	}
	conf := config.ReadConfig()
	_, err := git.PlainClone(conf.ProjectDirectory, false, &git.CloneOptions{
		URL:      URL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal("Failed to clone repository: ", err)
	}
}
