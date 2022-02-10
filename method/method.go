package main

import (
	"fmt"
	"reflect"
)

// type field struct {
// 	name string
// }

// func (p *field) print() {

// 	fmt.Println(p.name)
// }

// func main() {
// 	data1 := []*field{{"one"}, {"two"}, {"three"}}
// 	for _, v := range data1 {

// 		//go v.print()
// 		go (*field).print(v)
// 	}

// 	data2 := []field{{"four"}, {"five"}, {"six"}}
// 	for _, v := range data2 {

// 		//go v.print()
// 		fmt.Println(&v)
// 		go (*field).print(&v)
// 	}

// 	time.Sleep(3 * time.Second)
// }

// func main() {
// 	arr := [2]int{1, 2}
// 	res := []*int{}
// 	for _, v := range arr {
// 		fmt.Println(&v)
// 		res = append(res, &v)
// 	}
// 	//expect: 1 2
// 	fmt.Println(*res[0], *res[1])
// }

// type T struct {
// 	a int
// }

// func (t T) M1() {
// 	t.a = 10
// }

// func (t *T) M2() {
// 	t.a = 11
// }

type T struct{}

func (T) M1() {}
func (T) M2() {}

type S = T

func (*T) M3() {}
func (*T) M4() {}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty", dynTyp.Name())
		return
	}

	fmt.Printf("%s's method set:", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

func main() {
	// var t T
	// println(t.a)

	// t.M1()
	// println(t.a)

	// p := &t
	// p.M2()
	// println(t.a)

	// var n int
	// dumpMethodSet(n)
	// dumpMethodSet(&n)

	//var t T

	var s S

	dumpMethodSet(s)
	dumpMethodSet(&s)
}
