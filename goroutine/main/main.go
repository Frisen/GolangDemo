package main

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var balance int

func main() {
	for i := 0; i < 50; i++ {
		go SetBalance(i)
		go Balance(i)
	}
	time.Sleep(time.Second * 9)
}

func Balance(i int) int {
	//time.Sleep(time.Second)
	mu.RLock() // readers lock
	defer mu.RUnlock()
	//fmt.Println("balance")
	return balance
}
func SetBalance(i int) int {
	mu.Lock() // readers lock
	defer mu.Unlock()
	//time.Sleep(time.Second * 3)
	//fmt.Println("set")
	balance = balance + i
	return balance
}
