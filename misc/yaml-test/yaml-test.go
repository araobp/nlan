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

type T struct {
	A int
	B struct {
		RenamedC string `yaml:"c"`
		D        []int  `yaml:",flow"`
	}
}

func main() {

	m := make(map[interface{}]interface{})
	t := T{}

	_ = yaml.Unmarshal([]byte(data), &m)
	_ = yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- m:\n%v\n", m)
	fmt.Printf("--- m:\n%v\n", t)

	d, _ := yaml.Marshal(&m)
	d2, _ := yaml.Marshal(&t)
	fmt.Printf("--- m dump:\n%s\n", string(d))
	fmt.Printf("--- m dump:\n%s\n", string(d2))
}
