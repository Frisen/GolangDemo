package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	//fmt.Println(strings.Join(os.Args[0:], " "))
	//dup1()
	jsonDate()
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

//解析字符串
func jsonDate() {
	input := `{"created_at": "Thu May 31 00:00:01 +0000 2012"}`
	var dateMap map[string]interface{}
	if err := json.Unmarshal([]byte(input), &dateMap); err != nil {
		panic(err)
	}

	fmt.Println(dateMap)
	for k, v := range dateMap {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
