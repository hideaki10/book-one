package main

import "fmt"

// func setup(task string) func() {
// 	println("do some setup stuff for", task)
// 	return func() {
// 		println("do some teardown stuff for", task)
// 	}
// }

// func main() {
// 	teardown := setup("demo")
// 	defer teardown()
// 	println("do some bussiness stuff")
// }

// func greeting(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome, Gopher!\n")
// }
// func main() {
// 	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
// }

// func foo() {
// 	println("call foo")
// 	bar()
// 	println("end of foo")
// }

// func bar() {
// 	defer func() {
// 		if e := recover(); e != nil {
// 			fmt.Println("recovered the panic: ", e)
// 		}
// 	}()

// 	println("call bar")
// 	panic("panic occurs in bar")
// 	zoo()
// 	println("exit bar")

// }
// func zoo() {
// 	println("call zoo")
// 	println("exit zoo")
// }

// func main() {
// 	println("call main")
// 	foo()
// 	println("exit main")
// }

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	fmt.Println("foo1 result :")
	foo1()
	fmt.Println("\n foo2 result:")
	foo2()
	fmt.Println("\n foo3 result:")
	foo3()
}
