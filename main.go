package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["a"] = 1

	a := Path{
		{3, 4},
		{4, 2},
	}
	fmt.Print(a)
	b := []Point{
		{2, 3},
		{4, 2},
	}
	fmt.Println(b)
}

type Point struct {
	x float64
	y float64
}

type Path []Point
