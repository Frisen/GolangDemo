package main

import (
	"fmt"
	"testing"
)

func TestToM(t *testing.T) {
	a := "a_bcd_efg"
	r := ToM(a)
	fmt.Println(r)
}
