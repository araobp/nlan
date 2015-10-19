package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"
)

const filename = "roster.yaml"

func main() {

	v := make(map[string]string)
	v["router1"] = "192.168.56.101"
	v["router2"] = "192.168.56.102"
	v["router3"] = "192.168.56.103"

	roster_in, _ := ioutil.ReadFile(filename)
	templ, _ := template.New("roster").Parse(string(roster_in))
	var roster_out bytes.Buffer
	_ = templ.Execute(&roster_out, v)
	fmt.Println(roster_out.String())
}
