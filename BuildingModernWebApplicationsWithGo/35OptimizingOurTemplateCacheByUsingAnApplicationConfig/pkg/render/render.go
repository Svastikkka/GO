package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/config"
)

var app *config.AppConfig

// NewTemplate: Sets a config for template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, templ string) {

	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[templ]
	if !ok {
		log.Fatal("Template Not Found")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template into a browser")
	}
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Println("file path is wrong")
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
