package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/config"
	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/handlers"
	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/35OptimizingOurTemplateCacheByUsingAnApplicationConfig/pkg/render"
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

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(port, nil)
}
