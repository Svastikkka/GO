package main

import (
	"fmt"
	"net/http"

	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/30BuildingSimpleTemplateCache/pkg/handlers"
)

const port = ":8080"

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(port, nil)
}
