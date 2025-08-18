package template

import (
	"bytes"
	"text/template"
)

type HandlerTemplate struct {
	Name string
}

func (ht *HandlerTemplate) Handler() ([]byte, error) {

	temp, err := template.ParseFiles("handler.go.tmpl")
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

	temp, err := template.ParseFiles("handler_test.go.tmpl")
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
