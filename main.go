package main

import "fmt"

func main() {
	// 1
	// sx := []int{1, 2, 3, 4, 5}
	// // 2
	// sl := make([]byte, 6, 10)

	// // 3
	// sy := sx[2:5:5]

	// println(len(sx))
	// println(cap(sx))
	// println(len(sl))
	// println(cap(sl))
	// println(len(sy))
	// println(cap(sy))

	// u := [...]int{11, 12, 13, 14, 15}
	// fmt.Println("array:", u)
	// s := u[1:5] // [11, 12, 13, 14, 15]s := u[1:3]
	// fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	var sl1 []int
	var sl2 = []int{}

	fmt.Printf("sl1: %v\n", sl1)
	fmt.Printf("sl2: %v\n", sl2)
}
