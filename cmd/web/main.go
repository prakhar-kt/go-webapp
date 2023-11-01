package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/prakhar-kt/go-webapp/pkg/config"
	"github.com/prakhar-kt/go-webapp/pkg/handlers"
	"github.com/prakhar-kt/go-webapp/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	templateCache, err := render.CreateTemplateCache()

	if err != nil {

		log.Fatal("cannot create template cache")

	}

	app.TemplateCache = templateCache

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)

}
