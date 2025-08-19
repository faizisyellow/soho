package template

import (
	"bytes"
	"text/template"
)

type RouteTemplate struct {
	Name string
}

func (rt *RouteTemplate) NewRoute() ([]byte, error) {

	temp, err := template.ParseFiles("route.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "route", rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
