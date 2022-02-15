package main

// interface 1
// func main() {
// 	var a int64 = 13
// 	var i interface{} = a

// 	v1, ok := i.(int64)
// 	fmt.Printf("v1=%d,the type of v1 is %T,ok=%v", v1, v1, ok)
// 	fmt.Println()
// 	v2, ok := i.(string)
// 	fmt.Printf("v2=%s,the type of v2 is %T,ok=%v", v2, v2, ok)
// 	fmt.Println()
// 	v3 := i.(int64)
// 	fmt.Printf("v3=%d,the type of v3 is %T", v3, v3)
// 	fmt.Println()
// 	v4 := i.([]int)
// 	fmt.Printf("the type of v4 is %T\n", v4)
// }

type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}

func main() {
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}

	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}
