package channelstudy

import (
	"fmt"
	"reflect"
)

func Square(nums ...int) <-chan int {
	return square(nums...)
}

func square(nums ...int) <-chan int {
	slice := nums[:2]
	fmt.Println("slice:", slice)
	fmt.Println(reflect.TypeOf(nums))
	a := [...]int{4, 5, 6}
	fmt.Println("a type is : ", reflect.TypeOf(a).String())
	ch := make(<-chan int)
	return ch
}
