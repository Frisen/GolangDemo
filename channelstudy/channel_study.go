package channelstudy

import (
	"fmt"
	"strconv"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

//SumArr ...
func SumArr() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	x := <-c
	y := <-c
	fmt.Println(x, y)
}

//CakeStore ...
func CakeStore() {
	yelloChan := make(chan string)
	redChan := make(chan string)
	defer func() {
		fmt.Println("CakeStore defer runed...")
	}()
	go cakeMaker(yelloChan, 3, "yello")
	go cakeMaker(redChan, 4, "red")
	yelloEnd, redEnd := false, false
	go func() {
		for {
			if yelloEnd && redEnd {
				return
			}
			select {
			case cake, ok := <-yelloChan:
				if !ok {
					fmt.Println("yello closed ......")
					yelloEnd = true
				} else {
					fmt.Printf("recive cake : %s\n", cake)
				}
			case cake, ok := <-redChan:
				if !ok {
					fmt.Println("red closed ......")
					redEnd = true
				} else {
					fmt.Printf("recive cake : %s\n", cake)
				}

			}
		}
	}()
}

func cakeMaker(line chan string, count int, color string) {
	defer func() {
		if err := recover(); err != nil {
			s := fmt.Sprintf("panic----->%v", err)
			fmt.Println(s)
		}
	}()
	for i := 0; i < count; i++ {
		if i == 1 {
			panic("something happend!")
		}
		line <- color + " cake " + strconv.Itoa(i)
	}

	//close(line)
}
