package main

import "fmt"

var data = [8]int{
	24, 18, 12, 9, 17, 66, 32, 4,
}

func main() {

	var index int

	for k, v := range data {
		if v == 19 {
			index = k
			break
		}
		index = -1
	}

	fmt.Println(index)
}
