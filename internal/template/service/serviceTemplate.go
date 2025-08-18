package template

import (
	"bytes"
	"text/template"

	"github.com/faizisyellow/soho/internal/utils"
)

type ServiceTemplate struct {
	Name string
}

func (st *ServiceTemplate) NewService(modulePath string) ([]byte, error) {

	mod, err := utils.GetModuleName(modulePath)
	if err != nil {
		return nil, err
	}

	data := struct {
		Module string
		Name   string
	}{
		Module: mod,
		Name:   st.Name,
	}

	temp, err := template.ParseFiles("service.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "service", data)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (st *ServiceTemplate) ServiceTestTemplate() ([]byte, error) {

	temp, err := template.ParseFiles("service_test.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "serviceTest", st.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (st *ServiceTemplate) ServiceMapTemplate() ([]byte, error) {
	temp, err := template.ParseFiles("serviceMap.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "serviceMap", st.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}

func (st *ServiceTemplate) ServiceImplementation() ([]byte, error) {

	temp := template.New("serviceImplementation")

	temp = template.Must(temp.Parse("{{.}}Service{Repo: store}"))

	b := new(bytes.Buffer)

	err := temp.Execute(b, st.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
