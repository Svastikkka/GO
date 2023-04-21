package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 6, 4, 3, 2, 1, 3, 4}
	BubbleSort(elements)
	if elements[0] != 1 {
		t.Error("First element should be 9")
	}
	if elements[7] != 9 {
		t.Error("First element should be 1")
	}
}
