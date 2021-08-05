package tree_test

import (
	"GolangDemo/tree"
	"fmt"
	"testing"
)

func Test_Factorial(t *testing.T) {
	res := tree.Factorial(10, 10)
	fmt.Println(res)
}
