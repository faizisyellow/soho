package template

import (
	"bytes"
	"text/template"
)

type RepositoryTemplate struct {
	Name string
}

func (rt *RepositoryTemplate) NewRepository() ([]byte, error) {

	temp, err := template.ParseFiles("repository.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "repository", rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (rt *RepositoryTemplate) RepositoryTest() ([]byte, error) {

	temp, err := template.ParseFiles("repository_test.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "repositoryTest", rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (rt *RepositoryTemplate) RepositoryMap() ([]byte, error) {

	temp, err := template.ParseFiles("repositoryMap.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.ExecuteTemplate(b, "repositoryMap", rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (rt *RepositoryTemplate) RepositoryImplementation() ([]byte, error) {
	temp := template.New("repositoryImplementation")

	temp = template.Must(temp.Parse("{{.}}{Db: db}"))

	b := new(bytes.Buffer)

	err := temp.Execute(b, rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}
