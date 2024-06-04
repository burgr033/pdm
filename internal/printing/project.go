package printing

import (
	"os"

	"github.com/burgr033/pdm/internal/project"
	"github.com/olekukonko/tablewriter"
)

// PrintProjectTable prints the project table.
// It prints the project name, category, origin, and state.
// exported = yes
func PrintProjectTable() {
	projects := project.GetProjectTable()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Category", "Origin", "State"})
	for _, v := range projects {
		table.Append(v)
	}
	table.Render()
}
