package pipeline

import (
	"bytes"
	"html/template"

	"gopkg.in/yaml.v3"
)

func Parse(input []byte) (Pipeline, error) {
	var raw Pipeline

	err := yaml.Unmarshal(input, &raw)
	if err != nil {
		return raw, err
	}

	tmpl, err := template.New("SelfTemplate").Parse(string(input))
	if err != nil {
		return raw, err
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, raw)

	if err != nil {
		return raw, err
	}

	err = yaml.Unmarshal(tpl.Bytes(), &raw)

	if err != nil {
		return raw, err
	}

	return raw, err
}
