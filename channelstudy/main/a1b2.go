package main

import "fmt"

func main() {
	leter := "ABCDEFGHIJ"
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	leterChan, numberChan := make(chan bool), make(chan bool)
	done := make(chan bool)
	go func() {
		for i, v := range leter {
			<-leterChan
			fmt.Print(string(v))
			if i == len(leter)-1 {
				break
			}
			numberChan <- true
		}
		done <- true
	}()
	go func() {
		for _, v := range numbers {
			<-numberChan
			fmt.Print(v)
			leterChan <- true
		}
	}()
	numberChan <- true
	<-done
}
