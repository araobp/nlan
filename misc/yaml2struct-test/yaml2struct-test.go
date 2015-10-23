package main

import (
	"bytes"
	"fmt"
	"text/template"

	_ "encoding/json"
	mapstruct "github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

func Yaml2Struct(text *string, t interface{}, v map[string]interface{}) {

	m := make(map[string]interface{})
	var text_ *string
	if v != nil {
		var out bytes.Buffer
		templ, _ := template.New("temp").Parse(*text)
		_ = templ.Execute(&out, v)
		outstring := out.String()
		text_ = &outstring
	} else {
		text_ = text
	}
	_ = yaml.Unmarshal([]byte(*text_), &m)
	_ = mapstruct.Decode(m, t)
}

func main() {

	type T struct {
		A string
		B struct {
			C string
			D []string
		}
	}

	var text = `
a: {{.router1}}
b:
  c: This is a string
  d: [{{.router2}}, {{.router3}}]
`
	v := map[string]interface{}{
		"router1": "192.168.56.101",
		"router2": "192.168.56.102",
		"router3": "192.168.56.103",
	}

	var text2 = `
a: "10.0.0.1"
b:
  c: This is a string
  d: ["10.0.0.2", "10.0.0.3"]
`

	t := T{}

	Yaml2Struct(&text, &t, v)
	fmt.Println(t)
	Yaml2Struct(&text2, &t, nil)
	fmt.Println(t)
}
