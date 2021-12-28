package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
	Talk()
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func (stu *Stduent) Talk() {}

func Talkit(t People) {
	fmt.Println("aaaaa")
}

func main() {

	//var c chan interface{} = make(chan interface{}, 20)

	// var peo People = new(Stduent)
	// think := "love"
	// fmt.Println(peo.Speak(think))
	// a := []int(nil)
	a := make([]int, 4)
	a = append(a, 3, 4, 5)
	if a == nil {
		fmt.Println("a is nil")
	}
	fmt.Println(a)
}
