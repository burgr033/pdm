package main

import (
	c "github.com/burgr033/pdm/internal/config"
	p "github.com/burgr033/pdm/internal/projects"
)

func main() {
	cfg := c.ReadConfig()
	c.AddProjectToConfig(cfg, "Gude", "https://github.com/burgr033/pdm", "devProject")
	p.ListProjects(cfg.Projects)
}
