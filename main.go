package main

import (
	"bytes"
	"embed"
	"errors"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	//go:embed templates
	templateFiles embed.FS

	domain      string
	domainTitle string
	templates   = []string{
		"air",
		"gci",
		"gi",
		"service",
	}
)

func main() {

	// v, err := os.ReadDir("templates")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, j := range v {
	// 	k := j.Name()
	// 	log.Print(k)
	// }
	args := map[string]string{
		"Domain":      domain,
		"DomainTitle": domainTitle,
		"Date":        time.Now().Format(time.RFC3339),
	}
	//TODO: replace with actual project name.
	if err := os.Mkdir("project", 0755); err != nil {
		if errors.Is(err, os.ErrExist) {
			log.Print("exists")
		}
		log.Fatalf("failed to create %s directory %s", domain, err)
	}

	for _, tmplName := range templates {
		log.Printf("Creating %s.go file\n", tmplName)
		tmpl, err := template.ParseFS(templateFiles, "templates/"+tmplName+".tmpl")
		if err != nil {
			log.Fatalf("failed to read template %s.go.tmpl: %s", tmplName, err)
		}
		buf := bytes.Buffer{}
		if err := tmpl.Execute(&buf, args); err != nil {
			log.Fatalf("failed to parse template %s: %s", tmplName, err)
		}
		fn := tmplName + ".go"
		if tmplName == "domain" {
			fn = domain + "s.go"
		}
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatalf("go/format: %s", err)
		}
		if err := ioutil.WriteFile("../../"+domain+"s/"+fn, formatted, 0644); err != nil {
			log.Fatalf("failed to write template %s: %s", tmplName, err)
		}
		log.Printf("Created %s.go file successfully\n", tmplName)
	}
}
