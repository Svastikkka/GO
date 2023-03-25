package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func renderTemplate(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template")
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")

}

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(port, nil)
}
