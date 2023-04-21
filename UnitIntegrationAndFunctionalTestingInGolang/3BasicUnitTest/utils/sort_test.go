package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 6, 4, 3, 2, 1, 3, 4}
	BubbleSort(elements)
}
