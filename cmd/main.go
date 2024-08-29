package main

import (
	"fiber-generator/cmd/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var baseDir string

func readTemplateFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return string(content), nil
}

func generate(ctx *cli.Context) error {
	fmt.Println("generating")

	// Define base directories
	viewDir := filepath.Join(baseDir, "views")
	publicDir := filepath.Join(baseDir, "public")
	publicCSSDir := filepath.Join(publicDir, "css")
	publicJSDir := filepath.Join(publicDir, "js")
	publicLibDir := filepath.Join(publicDir, "lib")
	routesDir := filepath.Join(baseDir, "routes")
	controllersDir := filepath.Join(baseDir, "controllers")

	// Define template paths
	templatePaths := map[string]string{
		filepath.Join(viewDir, "index.html"):                  filepath.Join("cmd", "templates", "html", "index.html"),
		filepath.Join(publicCSSDir, "style.css"):              filepath.Join("cmd", "templates", "css", "style.css"),
		filepath.Join(publicJSDir, "app.js"):                  filepath.Join("cmd", "templates", "js", "index.js"),
		filepath.Join(publicLibDir, "htmx.min.js"):            filepath.Join("cmd", "templates", "lib", "htmx.min.js"), // Adjust as needed
		filepath.Join(routesDir, "userRoutes.go"):             filepath.Join("cmd", "templates", "go", "userRoutes.go"),
		filepath.Join(controllersDir, "CreateUserHandler.go"): filepath.Join("cmd", "templates", "go", "userHandler.go"),
		filepath.Join(baseDir, "app.go"):                      filepath.Join("cmd", "templates", "go", "app.go"),
	}

	// Create directories
	dirs := []string{viewDir, publicCSSDir, publicJSDir, publicLibDir, routesDir, controllersDir}
	for _, dir := range dirs {
		if err := utils.CreateDir(dir); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	// Create files from templates
	for filePath, templatePath := range templatePaths {
		content, err := readTemplateFile(templatePath)
		if err != nil {
			return err
		}
		if err := utils.CreateFile(filePath, content); err != nil {
			return fmt.Errorf("error creating file %s: %w", filePath, err)
		}
	}

	fmt.Println("done generating views")
	fmt.Println("done generating public")
	fmt.Println("done generating routes and controllers")

	return nil
}

func main() {
	app := &cli.App{
		Name:  "fiber-generator",
		Usage: "generate go fiber templates",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Value:       "fiberApp",
				Destination: &baseDir,
				Required:    true,
			},
		},
		Action: generate,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

