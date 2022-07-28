package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 1
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 1; i <= 4; i++ {
			<-ch1
			a++
			fmt.Println("协程一", a)
			ch2 <- true
		}
		<-ch1
		defer wg.Done()
	}()
	ch1 <- true
	go func() {
		for i := 1; i <= 4; i++ {
			<-ch2
			a++
			fmt.Println("协程二", a)
			ch1 <- true
		}
		defer wg.Done()
	}()
	wg.Wait()
	defer close(ch1)
	defer close(ch2)
}