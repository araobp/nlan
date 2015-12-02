package util

import (
	"bytes"
	"text/template"

	mapstruct "github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// This function converts a YAML text into a struct.
// yamltext is a YAML text that may include a place holder as a template.
// s must be a pointer to an empty struct.
// v is a dictionary for template.Execute
func Yaml2Struct(yamltext *string, s interface{}, v map[string]interface{}) {

	m := make(map[string]interface{})
	var yamltext_ *string
	if v != nil {
		var out bytes.Buffer
		templ, _ := template.New("temp").Parse(*yamltext)
		_ = templ.Execute(&out, v)
		outstring := out.String()
		yamltext_ = &outstring
	} else {
		yamltext_ = yamltext
	}
	_ = yaml.Unmarshal([]byte(*yamltext_), &m)
	_ = mapstruct.Decode(m, s)
}
