package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//  init
	elements := []int{9, 6, 4, 3, 2, 1, 3, 4}
	// execution
	BubbleSort(elements)
	// validation
	if elements[0] != 1 {
		t.Error("First element should be 1")
	}
	if elements[7] != 9 {
		t.Error("First element should be 9")
	}
}

// Below elements are already sorted but coverage is 100 % strange don't relay on coverage
func TestBubbleSort2(t *testing.T) {
	//  init
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// execution
	BubbleSort(elements)
	// validation
	if elements[0] != 1 {
		t.Error("First element should be 1")
	}
	if elements[7] != 8 {
		t.Error("First element should be 8")
	}
}
