package goroutine

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//RunWithMaxProd ...
func RunWithMaxProd() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	fmt.Println(n)
	time.Sleep(time.Second * 5)
	for {
		go fmt.Println("0")
		fmt.Println("1")
	}
}

//MutiSenderSingleReceipt 关闭channle
func MutiSenderSingleReceipt() {
	rand.Seed(time.Now().UnixNano())
	maxNum := 10000
	//发生者数量
	senderCount := 100
	//关闭chan
	stopCh := make(chan struct{})
	//Data chan
	dataCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < senderCount; i++ {
		go func() {
			for {
				value := rand.Intn(maxNum)
				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}

	go func() {
		defer wg.Done()
		var a struct{}
		for res := range dataCh {
			if res == maxNum-1 {
				//close(stopCh)
				stopCh <- a
				return
			}
			fmt.Println("number is ", res)
		}
	}()

	wg.Wait()
}
