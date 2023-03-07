package main


import (
   "fmt"
)

func Swap(sl []int, i int) {
	buff := sl[i]
	sl[i] = sl[i-1]
	sl[i-1] = buff
 }

func BubbleSort(nums []int) []int {
	prev := nums[0]
	count := 0
	unsorted := true
	for unsorted {
		count = 0
 
 
		for i, elem := range nums {
			if elem < prev && i > 0 {
 
 
				Swap(nums, i)
				count++
			}
			prev = nums[i]
		}
		if count == 0 {
			unsorted = false
		}
	}
 
 
	return nums
 }

func main() {
   var intSli []int
   var e rune
   fmt.Println("Please enter integers, or any letter to finish input:")
   for {
       e = 0
       fmt.Scan(&e)
       if e == 0 {
           fmt.Println("Input finished")
           break
       }
       intSli = append(intSli, int(e))
   }
   fmt.Println("Unsorted array:")
   fmt.Println(intSli)
   BubbleSort(intSli)
   fmt.Println("Sorted array:")
   for _, i := range intSli {fmt.Print(i," ")}
}