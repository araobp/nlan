package util

import (
	"strings"
)

// Python zfill-like function
func Zfill(s string, pad string, overall int) string {
	l := overall - len(s)
	return strings.Repeat(pad, l) + s
}

// Converts IP address into zero-padded string.
// For example, "10.1.2.3" => "010001002003"
func ZfillIp(ip string) string {
	splited := strings.Split(ip, ".")
	for i, s := range splited {
		splited[i] = Zfill(s, "0", 3)
	}
	return strings.Join(splited, "")
}
