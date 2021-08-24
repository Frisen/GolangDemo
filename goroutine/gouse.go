package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var balance int

func Try() {
	go SetBalance()
	go Balance()
	time.Sleep(time.Second * 9)
}

func Balance() int {
	time.Sleep(time.Second)
	mu.RLock() // readers lock
	defer mu.RUnlock()
	fmt.Println("balance")
	return balance
}
func SetBalance() int {
	// mu.Lock() // readers lock
	// defer mu.Unlock()
	time.Sleep(time.Second * 3)
	fmt.Println("set")

	return balance
}

/*func A(str string) {
	fmt.Println("A----0----", str)
	time.Sleep(time.Second)
	fmt.Println("A----1----", str)
	time.Sleep(time.Second)
	fmt.Println("A----2----", str)
}

func B(str string) {
	fmt.Println("B----", str)
	A(str)
}

func Exec() {
	for i := 0; i < 10; i++ {
		go B(fmt.Sprint(i))
	}
}

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	//for {
	cake := new(Cake)
	cake.state = "cooked"
	cooked <- cake // baker never touches this cake again
	//}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
		close(iced)
	}
}

func CakeMake() {
	cake := make(chan *Cake)
	icedcake := make(chan *Cake)
	go baker(cake)
	go icer(icedcake, cake)

	for v := range icedcake {
		fmt.Println(v.state)
	}
}
*/
