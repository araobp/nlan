package main

import (
	"fmt"
	"os"
)

type T string

func (s T) println() {
	fmt.Println(s)
}

func main() {
	message := os.Args[1]
	var s T = T(message)
	s.println()
}
