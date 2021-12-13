package sort

import (
	"fmt"
	"testing"
)

func TestQuiksort(t *testing.T) {
	// arr := []int{4, 2, 7, 5, 3, 6, 8, 1, 9}
	arr := []int{43, 99, 23, 56, 13, 4, 67, 87, 22, 73, 5, 66, 3, 6, 8, 1, 9}

	fmt.Println(arr)
	QuickSort(arr)
	fmt.Println(arr)
}
