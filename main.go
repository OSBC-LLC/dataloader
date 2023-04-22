package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	pkgstr "github.com/OSBC-LLC/dataloader/pkg/string"
	"github.com/joho/godotenv"
)

var port = "8080"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env err: ", err)
	} else {
		port = os.Getenv("PORT")
	}
}

func main() {
	maleFirst, _ := pkgstr.GetRandomMaleFirstName()
	maleLast, _ := pkgstr.GetRandomLastName()
	fmt.Println("Name: ", maleFirst, maleLast)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is for the root path only
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Load and parse all template files in the "templates" folder
	tmpl, err := loadTemplates("templates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and send it as the response
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadTemplates(templatesDir string) (*template.Template, error) {
	tmpl := template.New("")

	err := filepath.Walk(templatesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-HTML files and directories
		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		// Parse the template file and add it to the template set
		_, err = tmpl.ParseFiles(path)
		return err
	})

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
