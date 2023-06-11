package main

import (
	"fmt"
	"log"
	"net/http"

	"35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/config"
	"35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/handlers"
	"35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/render"
)

const port = ":8080"

func main() {
	fmt.Print("Starting server at port", port)
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	_ = http.ListenAndServe(port, nil)
}
