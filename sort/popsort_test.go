package sort

import (
	"fmt"
	"testing"
)

func TestPopsort(t *testing.T) {
	a := []int{43, 23, 98, 2, 8, 1, 15, 25, 88}
	Popsort(a)
	fmt.Println(a)
}
