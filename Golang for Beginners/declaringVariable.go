package main

import "fmt"

func main() {
	// Declare but value assign later
	var name string
	name = "manshu sharma"
	fmt.Println(name)

	// Short hand way
	var firstname, surname string = "masnhu", "sharma"
	fmt.Println(firstname, surname)

	// Short hand way for diff data type
	var (
		country string = "India"
		code    int    = +91
	)
	fmt.Println(country, code)

	// Short variable declaration or implicit declaration
	fav_language := "Python"
	fmt.Print(fav_language)
}
