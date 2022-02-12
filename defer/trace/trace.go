package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

var goroutineSpace = []byte("goroutine ")

func CurGoroutineID() uint64 {

	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]

	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')

	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}

	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func Trace() func() {

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := CurGoroutineID()
	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)

	return func() {
		fmt.Printf("g[%05d]: exit: [%s]\n", gid, name)
	}
}
