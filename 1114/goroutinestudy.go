package main

import (
	"log"
	"sync"
)

func main() {
	w := sync.WaitGroup{}
	w.Add(3)
	var f *Foo = &Foo{}
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	go func() {
		f.third(c2)
		w.Done()
	}()
	go func() {
		f.second(c1, c2)
		w.Done()
	}()
	go func() {
		f.first(c1)
		w.Done()
	}()
	w.Wait()
}

type Foo struct{}

func (*Foo) first(c1 chan interface{}) {
	log.Println("first")
	c1 <- 1
}

func (*Foo) second(c1, c2 chan interface{}) {
	<-c1
	log.Println("second")
	c2 <- 1
}

func (*Foo) third(c2 chan interface{}) {
	<-c2
	log.Println("third")
}
