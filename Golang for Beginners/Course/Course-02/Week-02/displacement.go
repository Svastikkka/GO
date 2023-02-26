package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, ivelocity, idisplacement float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (acceleration*(math.Pow(t, 2)*0.5) + (ivelocity * t) + (idisplacement))
	}
	return fn
}
func main() {
	fmt.Println("Please enter absolute file path")
	var acceleration, ivelocity, idisplacement, time float64
	fmt.Println("Enter Acceleration")
	fmt.Scan(&acceleration)
	fmt.Println("Enter Initial velocity")
	fmt.Scan(&ivelocity)
	fmt.Println("Enter Initial displacement")
	fmt.Scan(&idisplacement)
	fmt.Println("Enter Time")
	fmt.Scan(&time)
	fn := GenDisplaceFn(acceleration, ivelocity, idisplacement)
	fmt.Printf("Displacement: %f\n", fn(time))
}
