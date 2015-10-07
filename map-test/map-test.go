package main

import (
	"fmt"
)

func main() {

	// Python-dict-like map
	m := make(map[interface{}]interface{})

	m["key1"] = 1
	m["key2"] = "this is string"

	fmt.Println(m["key1"])
	fmt.Println(m["key2"])
}
