package main

import (
	"fmt"
)

func main() {
	// 	var sl = [][]int{{1, 34, 26, 35, 78}, {3, 45, 13, 24, 99}, {101, 13, 38, 7, 127}, {54, 27, 40, 83, 81}}

	// outloop:
	// 	for i := 0; i < len(sl); i++ {
	// 		for j := 0; j < len(sl[i]); j++ {
	// 			if sl[i][j] == 99 {
	// 				fmt.Printf("found 13 at [%d,%d]\n", i, j)
	// 				continue outloop
	// 			}
	// 		}
	// 	}

	// var m = []int{1, 2, 3, 4, 5}

	// for i, v := range m {
	// 	go func() {
	// 		time.Sleep(time.Second * 1)
	// 		fmt.Println(i, v)
	// 	}()
	// }

	// time.Sleep(time.Second * 2)

	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a = ", a)

	for i, v := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r = ", r)
	fmt.Println("after for range loop, a = ", a)
}
