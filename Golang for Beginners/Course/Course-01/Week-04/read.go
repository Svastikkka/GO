package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

const (
	ListSize = 10
)

func readfile(filename string) {
	// fmt.Println(filename)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("Error happen", err)
	}

	reader := bufio.NewReader(file)
	var personSlice = make([]Person, 0, ListSize)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Split(string(line), " ")
		person := Person{fname: fields[0], lname: fields[1]}
		personSlice = append(personSlice, person)
	}

	// fmt.Println(personSlice)
	for _, names := range personSlice {
		// Printing  the first and last names found in each struct
		fmt.Printf("First name: %-20s Last name: %-20s\n", names.fname, names.lname)

	}
}

func main() {
	fmt.Println("Please enter absolute file path")
	filename := bufio.NewScanner(os.Stdin)
	filename.Scan()
	readfile(filename.Text())
}
