package main

import (
	"GolangDemo/channelstudy"
	"bufio"
	"fmt"
	"os"
)

type person struct {
	name string
	age  int
}

func main() {
	channelstudy.CakeStore()

	// var v1 map[int]string
	// if v1 != nil {
	// 	fmt.Println("v1 is not nil -->", v1)
	// }
	// v2 := make(map[int]string, 2)
	// if v2 != nil {
	// 	fmt.Println("v2 is not nil,-->", v2)
	// }

	// var v3 chan string
	// if v3 != nil {
	// 	fmt.Println("v2 is not nil,-->", v3)
	// }

	// v4 := make(chan string)
	// go func() {
	// 	v4 <- "first ball"
	// }()
	// if v4 != nil {
	// 	fmt.Println("v4 is not nil,-->", <-v4)
	// }

	// v5 := make([]int, 2)
	// if v5 != nil {
	// 	fmt.Println("v5 is not nil,-->", v5)
	// }

	// v6 := new(int)
	// if v6 != nil {
	// 	*v6 = 12
	// 	fmt.Println("v6 is not nil,-->", *v6)
	// } else {
	// 	fmt.Println("v6 is nil,but now is -->", *v6)
	// }
	/*
		//example.JSONDate()
		var u dto.Users
		var ua dto.UsersA
		ip.GetEntity(nil, &u)
		ip.GetEntity(nil, &ua)
		fmt.Println(u, ua)
	*/
}

//从命令行获取输入的包bufio
//获取相同行
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}

		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
