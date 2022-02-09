package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		fmt.Println("data1")
		fmt.Println(&v)
		go v.print()
		//go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		fmt.Println("data2")
		fmt.Println(&v)
		go v.print()
		//go (*field).print(&v)
	}

	time.Sleep(3 * time.Second)
}
