package projects

import (
	"fmt"

	c "github.com/burgr033/pdm/internal/config"
)

func ListProjects(ProjectList []c.Project) {
	for v := range len(ProjectList) {
		fmt.Printf("[%v]: %v\n", ProjectList[v].ID, ProjectList[v].Name)
	}
}
