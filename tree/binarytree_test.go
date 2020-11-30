package tree_test

import (
	"GolangDemo/tree"
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	elements := [...]int{3, 6, 8, 4, 2}
	bst := tree.Add(5, nil)
	for _, v := range elements {
		tree.Add(v, bst)
	}
	fmt.Println("length:", bst.Length)
	res := tree.Contains(bst, 3)
	fmt.Println("Is Contains:", res)
}

func Test_PreOrder(t *testing.T) {
	elements := []interface{}{3, 6, 8, 4, 2}
	bst := tree.Add(5, nil)
	for _, v := range elements {
		tree.Add(v, bst)
	}
	tree.PreOrder(bst)
}
