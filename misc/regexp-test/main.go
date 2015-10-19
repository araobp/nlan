package main

import (
	"fmt"
	"regexp"
)

var text string = `- <router>
- <router>`

func main() {

	re := regexp.MustCompile("<.*>")
	fmt.Println(re.ReplaceAllString(text, "10.56.41.145"))
}
