package category

import (
	"log"
	"os"
	"path"

	"github.com/burgr033/pdm/internal/config"
)

// AddCategory adds a new category to the config file.
// It takes the name and relative path of the category as arguments.
// It reads the config file to get the list of categories.
// If the category already exists, it logs a message and exits the function.
// Otherwise, it appends the new category to the config.Categories slice and calls addCategoryDirectory.
// exported = yes
func AddCategory(name string, relPath string) {
	conf := config.ReadConfig()
	if DoesCategoryAlreadyExist(name, relPath) {
		log.Printf("Category already exists: %v", name)
	} else {
		conf.Categories = append(conf.Categories, config.Category{Name: name, RelPath: relPath})
		addCategoryDirectory(relPath)
		config.WriteConfig(conf)
	}
}

// AddCategory adds a new category to the config file.
// It checks if the category already exists by calling DoesCategoryAlreadyExist.
// If the category does not exist, it appends the new category to the config.Categories slice.
// It then calls addCategoryDirectory to create the category directory.
// Finally, it writes the updated config to the config file.
// exported = no
func addCategoryDirectory(relPath string) {
	conf := config.ReadConfig()
	fullPath := path.Join(conf.ProjectDirectory, relPath)
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create category directory: %v", err)
	}
}

// addCategoryDirectory creates a new directory for the category.
// It reads the config file to get the project directory path.
// It then joins the project directory path with the relative path of the category.
// It creates the directory with the full path and sets the permissions to 0755.
// If an error occurs, it logs the error and exits the program.
// exported = no
func doesCategoryNameExist(name string) bool {
	conf := config.ReadConfig()
	for _, c := range conf.Categories {
		if c.Name == name {
			return true
		}
	}
	return false
}

// doesCategoryNameExist checks if a category with the same name already exists.
// It reads the config file to get the list of categories.
// It then iterates over the categories and checks if the name matches the given name.
// If a category with the same name exists, it returns true.
// Otherwise, it returns false.
// exported = yes
func DoesCategoryDirectoryExist(relPath string) bool {
	conf := config.ReadConfig()
	for _, c := range conf.Categories {
		if c.RelPath == relPath {
			fullPath := path.Join(conf.ProjectDirectory, c.RelPath)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				return false
			}
			return true
		}
	}
	return false
}

// DoesCategoryDirectoryExist checks if the directory for a category already exists.
// It reads the config file to get the list of categories.
// It then iterates over the categories and checks if the relative path matches the given relative path.
// If the directory exists, it returns true.
// Otherwise, it returns false.
// exported = yes
func DoesCategoryAlreadyExist(name string, relPath string) bool {
	return doesCategoryNameExist(name) || DoesCategoryDirectoryExist(relPath)
}

// GetCategoryList returns a list of category names.
// It reads the config file to get the list of categories.
// It then iterates over the categories and appends the names to a slice.
// Finally, it returns the slice of category names.
// exported = yes
func GetCategoryList() []string {
	conf := config.ReadConfig()
	var categories []string
	for _, c := range conf.Categories {
		categories = append(categories, c.Name)
	}
	return categories
}

// GetCategoryTable returns a table of category names and relative paths.
// It reads the config file to get the list of categories.
// It then iterates over the categories and appends the name and relative path to a slice.
// Finally, it returns the slice of category names and relative paths.
// exported = yes
func GetCategoryTable() [][]string {
	conf := config.ReadConfig()
	var categories [][]string
	for _, c := range conf.Categories {
		categories = append(categories, []string{c.Name, c.RelPath})
	}
	return categories
}

// CountProjectsInCategory returns the number of projects in a category.
// It takes the name of the category as an argument.
// It gets the full path of the category by calling GetFullCategoryPath.
// It reads the directory and counts the number of files.
// Finally, it returns the number of projects in the category.
// exported = yes
func CountProjectsInCategory(name string) int {
	fullPath := GetFullCategoryPath(name)
	files, err := os.ReadDir(fullPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}
	return len(files)
}

// GetFullCategoryPath returns the full path of a category.
// It takes the name of the category as an argument.
// It reads the config file to get the list of categories.
// It then iterates over the categories and checks if the name matches the given name.
// If the name matches, it returns the full path of the category.
// If no category with the given name is found, it return an error
// exported = yes
func GetFullCategoryPath(name string) string {
	conf := config.ReadConfig()
	for _, c := range conf.Categories {
		if c.Name == name {
			return path.Join(conf.ProjectDirectory, c.RelPath)
		}
	}
	return "404 not found"
}
