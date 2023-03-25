package main

import (
	"fmt"
	"github.com/svastikkka/BuildingModernWebApplicationsWithGo/pkg/handlers"
	"net/http"
)

const port = ":8080"

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(port, nil)
}
