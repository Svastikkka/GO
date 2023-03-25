package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(port, nil)
}
