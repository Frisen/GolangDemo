package main

import (
	"GolangDemo/example"
	"bufio"
	"fmt"
	"os"
)

func main() {
	example.JSONDate()
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
