package main

import (
	"fmt"
)

type I interface{}

func main() {
	var i I
	i = 3
	fmt.Print(i)
	i = "hello"
	fmt.Print(i)
}
