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
	// integer
	var unsigned_number1 uint8 = 111
	var unsigned_number2 uint16 = 11111
	var unsigned_number3 uint32 = 1111111111
	var unsigned_number4 uint64 = 1111111111111111111
	var unsigned_number5 uint = 1111111111111111111
	fmt.Println("UNSIGNED", unsigned_number1, unsigned_number2, unsigned_number3, unsigned_number4, unsigned_number5)
	// array
	// slices (same as array but provide more control over memory allocation)
	// map
}
