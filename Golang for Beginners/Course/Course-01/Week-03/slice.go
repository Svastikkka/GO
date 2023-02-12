package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Find: Finds a integer inside a slice
func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	var empty_slice = make([]int, 3, 3)
	var input string
	for {
		fmt.Scan(&input)
		// string to int
		i, err := strconv.Atoi(input)
		if input != "X" && err == nil {
			index, found := Find(empty_slice, 0)
			if !found {
				empty_slice = append(empty_slice, i)
			} else {
				empty_slice[index] = i
			}
			// Using Ints function
			sort.Ints(empty_slice)
			fmt.Println(empty_slice)
		} else {
			// handling err
			if input != "X" && err != nil {
				panic(err)
			}
			break
		}
	}
}
