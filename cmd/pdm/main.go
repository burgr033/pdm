package main

import (
	"log"
	"os"

	"github.com/burgr033/pdm/internal/category"
	"github.com/burgr033/pdm/internal/printing"
	"github.com/burgr033/pdm/internal/project"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "category",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a category",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Required: true,
								Name:     "name",
								Usage:    "name of the category",
							},
							&cli.StringFlag{
								Name:  "Path",
								Usage: "relative directory name of the category",
							},
						},
						Action: func(c *cli.Context) error {
							usedPath := ""
							if c.String("path") != "" {
								usedPath = c.String("path")
							} else {
								usedPath = c.String("name")
							}
							category.AddCategory(c.String("name"), usedPath)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list all categories",
						Action: func(c *cli.Context) error {
							printing.PrintCategoryTable()
							return nil
						},
					},
				},
			},
			{
				Name: "project",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list all projects",
						Action: func(c *cli.Context) error {
							printing.PrintProjectTable()
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a project",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Required: true,
								Name:     "name",
							},
							&cli.StringFlag{
								Required: true,
								Name:     "category",
							},
							&cli.StringFlag{
								Name: "url",
							},
						},
						Action: func(c *cli.Context) error {
							if c.String("url") != "" {
								project.CloneProject(c.String("url"), c.String("category"))
							} else {
								project.InitProject(c.String("name"), c.String("category"))
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
