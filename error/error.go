package main

import (
	"errors"
	"io"
	"strings"
)

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	err := errors.New("dasdasd")

// 	res := fmt.Errorf("dsadas", err)

// }

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}

	return p
}

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

func main() {
	r := strings.NewReader("hello,gopher\n")
	lr := io.LimitReader(r, 4)
	println(lr)
	printNilInterface()
	// err := returnsError()

	// if err != nil {
	// 	fmt.Printf("error occur: %+v\n", err)
	// 	return

	// }

	// fmt.Println("ok")
}
