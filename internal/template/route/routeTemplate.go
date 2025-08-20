package template

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
)

type RouteTemplate struct {
	Name string
}

type Router struct {
	Endpoint string
	Handler  string
}

//go:embed route.go.tmpl
var routeEmbed embed.FS

func (rt *RouteTemplate) NewRoute() ([]byte, error) {

	temp, err := template.ParseFS(routeEmbed, "route.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	data := Router{
		Endpoint: strings.ToLower(rt.Name),
		Handler:  rt.Name,
	}

	err = temp.ExecuteTemplate(b, "route", data)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
