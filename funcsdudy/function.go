package funcsdudy

import (
	"fmt"
)

//AddFunc ...
type AddFunc func() int

func getAddFunc() AddFunc {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Exec ...
func Exec() {
	myAddFunc := getAddFunc()
	fmt.Println(myAddFunc())
	fmt.Println(myAddFunc())
}
