package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-server/internals/config"
	"go-server/internals/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":3000"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
