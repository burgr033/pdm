package printing

import (
	"os"
	"strconv"

	"github.com/burgr033/pdm/internal/category"
	"github.com/olekukonko/tablewriter"
)

// PrintCategoryTable prints the category table.
// It prints the category name, path, project count, and if the category directory exists.
// exported = yes
func PrintCategoryTable() {
	categories := category.GetCategoryTable()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Path", "ProjectCount", "Exists"})
	for _, v := range categories {
		if category.DoesCategoryDirectoryExist(v[0]) {
			projectCount := category.CountProjectsInCategory(v[0])
			v = append(v, strconv.Itoa(projectCount))
			v = append(v, "true")
		} else {
			v = append(v, "0")
			v = append(v, "false")
		}
		table.Append(v)
	}
	table.Render()
}
