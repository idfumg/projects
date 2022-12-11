package yamltohtml

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PageData struct {
	Title string `yaml:"Title"`
	Body  string `yaml:"Body"`
}

func YamlToHtml(path string) (string, error) {
	t, err := template.New("page").Parse(`<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>`)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var pageData PageData
	err = yaml.Unmarshal(data, &pageData)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, pageData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
