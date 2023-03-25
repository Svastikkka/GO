package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func add(x, y int) int {
	return x + y
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := add(7, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a about page %d", sum))
}

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(port, nil)
}
