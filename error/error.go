package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("dasdasd")

	res := fmt.Errorf("dsadas", err)

}
