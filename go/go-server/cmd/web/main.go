package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/m3rashid/learn_x/go/go-server/pkg/config"
	"github.com/m3rashid/learn_x/go/go-server/pkg/handlers"
	"github.com/m3rashid/learn_x/go/go-server/pkg/render"
)

const portNumber = ":5000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
