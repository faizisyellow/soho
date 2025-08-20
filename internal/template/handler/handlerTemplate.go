package template

import (
	"bytes"
	"embed"
	"text/template"
)

type HandlerTemplate struct {
	Name string
}

var (

	//go:embed handler.go.tmpl
	handlerEmbed embed.FS

	//go:embed handler_test.go.tmpl
	handlerTestEmbed embed.FS
)

func (ht *HandlerTemplate) Handler() ([]byte, error) {

	temp, err := template.ParseFS(handlerEmbed, "handler.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "handler", ht.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (ht *HandlerTemplate) HandlerTest() ([]byte, error) {

	temp, err := template.ParseFS(handlerTestEmbed, "handler_test.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "handlerTest", ht.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
