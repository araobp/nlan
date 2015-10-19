package main

import (
	"bufio"
	"bytes"
	"fmt"
	_ "io/ioutil"
	"os"
	"text/template"
)

const filename = "roster.yaml"

func main() {

	v := make(map[string]string)
	v["router1"] = "192.168.56.101"
	v["router2"] = "192.168.56.102"
	v["router3"] = "192.168.56.103"

	//roster_in, _ := ioutil.ReadFile(filename)
	fp, _ := os.Open(filename)
	defer fp.Close()
	reader := bufio.NewReader(fp)
	head, _, _ := reader.ReadLine()
	fmt.Println("--- 1st line ---")
	fmt.Println(string(head))
	fmt.Println("----------------")
	fp.Seek(0, 0)
	reader = bufio.NewReader(fp)
	roster_in := make([]byte, 4096)
	_, _ = reader.Read(roster_in)
	templ, _ := template.New("roster").Parse(string(roster_in))
	var roster_out bytes.Buffer
	_ = templ.Execute(&roster_out, v)
	fmt.Println(roster_out.String())
}
