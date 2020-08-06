package goroutinestudy

import (
	"fmt"
	"sync"
)

//Sum ...
func Sum() {
	arr := make([]int, 3)
	for i := 0; i < 5; i++ {
		go addEle(arr, i)
	}
}

func addEle(arr []int, i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("add fail")
		} else {
			fmt.Println(i)
		}
	}()
	arr[i] = i + 1
}

/*Cal goroutine同步channel通信
 */
func Cal() {
	myChan := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go add(i, i+1, myChan)
	}
	for j := 0; j < 10; j++ {
		<-myChan
	}
}

func add(i, j int, myChan chan bool) {
	fmt.Printf("%d + %d = %d\n", i, j, i+j)
	myChan <- true
}

//Cal1 使用sync实现同步
func Cal1() {
	var myWaitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		myWaitGroup.Add(1)
		go add1(i, i+1, &myWaitGroup)
	}
	myWaitGroup.Wait()
}

func add1(i, j int, group *sync.WaitGroup) {
	fmt.Printf("%d + %d = %d\n", i, j, i+j)
	group.Done()
}
