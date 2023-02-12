package main

import "fmt"

func main() {
	// boolean
	var YES bool = true
	var NO bool = false
	fmt.Println(YES, NO)

	// character
	var A = 'a'
	fmt.Println(A)

	// string
	var name string = "Manshu Sharma"
	var DOB string = "21/02/1997"
	fmt.Println(name, "\n", DOB)

	// float
	var float_number1 float32 = 1.11111111111111111111111111111111111111111111111111111111111111111
	var float_number2 float64 = 1.11111111111111111111111111111111111111111111111111111111111111111
	fmt.Println(float_number1, float_number2)

	// integer
	var number1 int8 = 111
	var number2 int16 = 11111
	var number3 int32 = 1111111111
	var number4 int64 = 1111111111111111111
	var number5 int = 1111111111111111111
	fmt.Println(number1, number2, number3, number4, number5)

	// unsigned integer
	var unsigned_number1 uint8 = 111
	var unsigned_number2 uint16 = 11111
	var unsigned_number3 uint32 = 1111111111
	var unsigned_number4 uint64 = 1111111111111111111
	var unsigned_number5 uint = 1111111111111111111
	fmt.Println("UNSIGNED", unsigned_number1, unsigned_number2, unsigned_number3, unsigned_number4, unsigned_number5)

	// array

	// slices (same as array but provide more control over memory allocation)
	// Reference: https://www.geeksforgeeks.org/slices-in-golang/
	// 1. Using slice literal
	var slice1 = []string{"Geeks", "for", "geeks"}
	fmt.Println(slice1)

	// 2. Using an Array
	var arr1 = [4]string{"Geeks", "for", "Geeks", "GFG"}
	var slice2 = arr1[0:2]
	fmt.Println(slice2)

	// 3. Using already Existing Slice
	var original_slice = []string{"Geeks", "for", "Geeks", "GFG"}
	var slice3 = original_slice[0:2]
	var slice4 = original_slice[2:4]
	fmt.Println(slice3, slice4)

	// map
	// struct (aggregate data type)
}
