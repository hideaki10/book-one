package main

import (
	"sync"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func main() {

	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()
	go func() {
		consume(ch)
		wg.Done()
	}()
	wg.Wait()
	// ch1 := make(chan int)

	// go func() {
	// 	ch1 <- 13
	// }()

	// n := <-ch1

	// println(n)

	// ch1 := make(chan<- int, 1)
	// ch2 := make(<-chan int, 1)

	// <-ch1
	// ch2 <- 13
}
