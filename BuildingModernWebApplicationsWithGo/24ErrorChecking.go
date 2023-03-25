package main

import (
	"errors"
	"fmt"
	"net/http"
)

const port = ":8080"

func add(x, y int) int {
	return x + y
}
func divide(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannnot divide by 0 or lesser")
		return 0, err
	}
	return x / y, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := add(7, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a about page %d", sum))
}
func Divide(w http.ResponseWriter, r *http.Request) {
	divide, err := divide(7, 0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a divide page %f", divide))
}
func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	_ = http.ListenAndServe(port, nil)
}
