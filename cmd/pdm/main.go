package main

import (
	"github.com/burgr033/pdm/internal/git"
)

// main is the entry point of the pdm command line tool.
func main() {
	git.CloneAction("https://github.com/burgr033/asdf")
	git.CloneAction("https://github.com/burgr033/dotfiles")
	git.CloneAction("asdf")
}
