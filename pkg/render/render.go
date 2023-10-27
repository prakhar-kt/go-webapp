package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, templateName string) {

	var template *template.Template
	var err error

	// check if the template already exists in the cache

	_, inMap := templateCache[templateName]

	if !inMap {

		// create the template

		log.Println("Creatinf template and adding to the cache")

		err = createTemplateCache(templateName)

		if err != nil {
			log.Println(err)
		}

	} else {

		log.Println("using existing template from cache")

	}

	template = templateCache[templateName]

	err = template.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(templateName string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err

	}

	templateCache[templateName] = tmpl

	return nil

}
