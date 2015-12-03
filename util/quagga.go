package util

import "strings"

const C = "-c" // vtysh -c option

var CONF_T = "configure terminal"
var END = "end"
var WRITE_FILE = "write file"

func appendScript(args *[]string, arg string) {
	*args = append(*args, C)
	*args = append(*args, arg)
}

func VtyshBatch(script [][]string) []string {
	var args []string
	appendScript(&args, CONF_T)
	for _, arg := range script {
		appendScript(&args, strings.Join(arg, " "))
	}
	appendScript(&args, END)
	appendScript(&args, WRITE_FILE)
	return args
}
