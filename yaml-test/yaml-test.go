package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
a: 1
b:
  c: This is string
  d: [2, 3]
`

func main() {

	m := make(map[interface{}]interface{})

	_ = yaml.Unmarshal([]byte(data), &m)
	fmt.Printf("--- m:\n%v\n", m)

	d, _ := yaml.Marshal(&m)
	fmt.Printf("--- m dump:\n%s\n", string(d))
}
