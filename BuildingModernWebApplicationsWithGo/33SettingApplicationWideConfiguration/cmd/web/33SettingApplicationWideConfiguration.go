package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/33SettingApplicationWideConfiguration/pkg/handlers"
	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/33SettingApplicationWideConfiguration/pkg/render"
	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/33SettingApplicationWideConfiguration/pkg/config"
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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(port, nil)
}
