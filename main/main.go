package main

import (
	datastruct "GolangDemo/data_struct"
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type person struct {
	name string
	age  int
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello man")
}

type Person struct {
	AgeYears int
	Name     string
	Friends  []Person
}

var mode = flag.String("mode", "", "process mode")

func main() {
	fmt.Print("Enter text: \n")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)

	charMap := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	count := len([]rune(input))
	stack := datastruct.NewStack()
	for i := 0; i < count; i++ {
		if stack.Length() > 0 {
			if v, ok := charMap[stack.Peek().(byte)]; ok && v == input[i] {
				stack.Pop()
				continue
			}
		}
		stack.Push(input[i])
	}
	if stack.Length() > 0 {
		fmt.Println("not matched!")
	} else {
		fmt.Println("matched!")
	}
	return
	stop := make(chan interface{})
	//ch := make(chan int, 3)
	var ch chan int

	ep := 1
	fmt.Println(ep)
	//close(ch)

	go func() {
		<-ch
		fmt.Println("888888888")
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println("chan-->", v)
			default:
				fmt.Println("非阻塞")
			}
		}

	}()
	fmt.Println("lasjfajsflaj")
	<-stop
	//goroutine.RunWithMaxProd()

	//http.HandlerFunc("/", sayHi)
	//http.ListenAndserve(":8080", nil)
	// http.HandleFunc("/", sayHi)
	// http.ListenAndServe(":8080", nil)
	/* 阻塞主线程 I fixed a bug here!
	mychan := make(chan int)
	go func() {
		for {
			mychan <- 2
		}
	}()
	go func() {
		for {
			fmt.Println(<-mychan)
		}
	}()
	select {}
	*/

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
