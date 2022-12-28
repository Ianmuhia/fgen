package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"
)

var (
	//go:embed templates/*
	templateFiles embed.FS

	domain      string
	domainTitle string
	templates   = []string{
		"air",
		"contributors",
		"dockerignore",
		"gci",
		"gi",
		"license",
		"makefile",
		"readme",
	}
)

func main() {

	// args := map[string]string{
	// 	"Domain":      domain,
	// 	"DomainTitle": domainTitle,
	// 	"Date":        time.Now().Format(time.RFC3339),
	// }
	//TODO: replace with actual project name.
	if err := os.MkdirAll(filepath.Join("project"), os.ModePerm); err != nil {
		log.Fatalf("failed to create %s directory %s", domain, err)
	}
	if err := os.Chdir("project"); err != nil {
		log.Fatalf("could not switch dir , %v", err)
	}
	for _, tmplName := range templates {
		switch tmplName {
		case "air":
			createDotFiles("air.tmpl", ".air.toml")
		case "contributors":
			createDotFiles("contributors.tmpl", "CONTRIBUTORS.txt")
		case "dockerignore":
			createDotFiles("dockerignore.tmpl", ".dockerignore")
		case "gci":
			createDotFiles("gci.tmpl", ".golangci.yml")
		case "gi":
			createDotFiles("gi.tmpl", ".gitignore")
		case "license":
			createDotFiles("license.tmpl", "LICENSE")
		case "makefile":
			createDotFiles("makefile.tmpl", "Makefile")
		case "readme":
			createDotFiles("readme.tmpl", "README")
		default:
			log.Print("no file to create")
		}

		log.Printf("Created %s file successfully\n", tmplName)
	}
}

func createDotFiles(tmplName, actual string) {
	log.Printf("Creating %s file\n", tmplName)

	b, err := templateFiles.ReadFile("templates/" + tmplName)
	if err != nil {
		log.Fatalf("could not read template file , %v", err)
	}
	f, err := os.Create(filepath.Join(actual))
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	if err != nil {
		log.Fatalf("could not create file , %v", err)
	}
	if err := os.WriteFile(f.Name(), b, 0644); err != nil {
		log.Fatalf("failed to write template %s: %s", actual, err)
	}
}
