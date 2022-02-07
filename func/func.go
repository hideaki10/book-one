package main

import (
	"fmt"
	"net/http"
)

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

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
