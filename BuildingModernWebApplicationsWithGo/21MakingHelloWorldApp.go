package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("hello world")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			log.Fatal("Something bad happen")
		}
		fmt.Println(fmt.Sprintf("Bytes written %d", n))
	})
	_ = http.ListenAndServe(":8080", nil)
}
