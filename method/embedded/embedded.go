package embedded

import "io"

type MyInt int

type T struct {
	A int
	B int
}

type S struct {
	*MyInt
	T
	io.Reader
	S string
	N int
}

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}
