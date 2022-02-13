package main

import (
	"bookstore/defer/trace"
	"sync"
)

// defer 1
// func Trace(name string) func() {
// 	println("enter:", name)
// 	return func() {
// 		println("exit:", name)
// 	}
// }

// func foo() {
// 	defer Trace("foo")()
// 	bar()
// }

// func bar() {
// 	defer Trace("bar")()

// }
// func main() {
// 	defer Trace("main")()
// 	foo()
// }

// defer2
// func Trace() func() {

// 	pc, _, _, ok := runtime.Caller(1)
// 	if !ok {
// 		panic("not found caller")
// 	}

// 	fn := runtime.FuncForPC(pc)
// 	name := fn.Name()

// 	gid := trace.CurGoroutineID()
// 	fmt.Printf("goroutine %d: %s\n", gid, name)
// 	println("enter: ", name)
// 	return func() {
// 		println("goroutine %d: %s\n", gid, name)
// 	}
// }

// func foo() {
// 	defer Trace()()
// 	bar()
// }

// func bar() {
// 	defer Trace()()
// }

// func main() {
// 	defer Trace()()
// 	foo()
// }

// defer3

// func Trace() func() {

// 	pc, _, _, ok := runtime.Caller(1)
// 	if !ok {
// 		panic("not found caller")
// 	}

// 	fn := runtime.FuncForPC(pc)
// 	name := fn.Name()

// 	gid := trace.CurGoroutineID()
// 	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)

// 	return func() {
// 		fmt.Printf("g[%05d]: exit: [%s]\n", gid, name)
// 	}
// }
// func A1() {
// 	defer Trace()()
// 	B1()
// }

// func B1() {
// 	defer Trace()()
// 	C1()
// }

// func C1() {
// 	defer Trace()()
// 	D()
// }

// func D() {
// 	defer Trace()()
// }

// func A2() {
// 	defer Trace()()
// 	B2()
// }

// func B2() {
// 	defer Trace()()
// 	C2()
// }

// func C2() {
// 	defer Trace()()
// 	D()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		A2()
// 		wg.Done()
// 	}()

// 	A1()
// 	wg.Wait()
// }
//defer4

func A1() {
	defer trace.Trace()()
	B1()
}

func B1() {
	defer trace.Trace()()
	C1()
}

func C1() {
	defer trace.Trace()()
	D()
}

func D() {
	defer trace.Trace()()
}

func A2() {
	defer trace.Trace()()
	B2()
}

func B2() {
	defer trace.Trace()()
	C2()
}

func C2() {
	defer trace.Trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}
