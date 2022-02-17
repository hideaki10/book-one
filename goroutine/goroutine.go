package main

import (
	"fmt"
	"runtime"
	"time"
)

// go fmt.Println("I am a goroutine!")

// var c = make(chan int)
// go func(a,b int){
// 	c <- a + b
// }(3,4)

// c := srv.newConn(rw)
// go c.serve(connCtx)

// func spawn(f func() error) <-chan error {
// 	c := make(chan error)

// 	go func() {
// 		c <- f()
// 	}()

// 	return c

// }

// func main() {
// 	c := spawn(func() error {
// 		time.Sleep(2 * time.Second)
// 		return errors.New("timeout")
// 	})
// 	fmt.Println(<-c)
// }

func deadloop() {
	for {
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

var sequence int
