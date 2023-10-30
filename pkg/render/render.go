package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/prakhar-kt/go-webapp/pkg/config"

	"github.com/prakhar-kt/go-webapp/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {

		// create a template cache
		templateCache = app.TemplateCache

	} else {
		templateCache, _ = CreateTemplateCache()

	}

	// get the requested template from the cache

	template, ok := templateCache[tmpl]

	if !ok {
		log.Fatal("Could  not get template from the cache")
	}

	// create a new buffer to store the template execution
	buf := new(bytes.Buffer)

	// store the template in the buffer

	templateData = AddDefaultData(templateData)

	err := template.Execute(buf, templateData)

	if err != nil {

		log.Println(err)

	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// create an empty templates Cache
	tempCache := map[string]*template.Template{}

	// Get all the pages matching the *.pages.tmpl pattern
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return tempCache, err
	}

	// loop over all the pages and create templates
	for _, page := range pages {

		// Get the page name
		name := filepath.Base(page)

		// Parse the templates from the page and assign it the name from above
		templates, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}

		// Get the layouts
		layouts, err := filepath.Glob("./templates/*layout.tmpl")

		if err != nil {
			return tempCache, err
		}

		if len(layouts) > 0 {
			templates, err = templates.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}

		tempCache[name] = templates

	}

	return tempCache, nil

}
