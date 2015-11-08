package main

import (
	"fmt"
	"strings"
)

func Zfill(s string, pad string, overall int) string {
	l := overall - len(s)
	return strings.Repeat(pad, l) + s
}

func ZfillIpAddress(ip string) string {
	splited := strings.Split(ip, ".")
	for i, s := range splited {
		splited[i] = Zfill(s, "0", 3)
	}
	return strings.Join(splited, "")
}

func main() {
	fmt.Println(Zfill("1", "0", 3))
	fmt.Println(Zfill("12", "0", 3))
	fmt.Println(Zfill("123", "0", 3))
	fmt.Println(ZfillIpAddress("10.2.34.1"))
}
