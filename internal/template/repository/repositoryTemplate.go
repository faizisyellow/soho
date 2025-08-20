package template

import (
	"bytes"
	"embed"
	"text/template"
)

type RepositoryTemplate struct {
	Name string
}

var (
	//go:embed repository.go.tmpl
	repositoryEmbed embed.FS

	//go:embed repository_test.go.tmpl
	repositoryTestEmbed embed.FS

	//go:embed repositoryMap.go.tmpl
	repositoryMapEmbed embed.FS
)

func (rt *RepositoryTemplate) NewRepository() ([]byte, error) {

	temp, err := template.ParseFS(repositoryEmbed, "repository.go.tmpl")
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

	temp, err := template.ParseFS(repositoryTestEmbed, "repository_test.go.tmpl")
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

	temp, err := template.ParseFS(repositoryMapEmbed, "repositoryMap.go.tmpl")
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

	temp = template.Must(temp.Parse("{{.}}: &{{.}}Repository{Db: db},"))

	b := new(bytes.Buffer)

	err := temp.Execute(b, rt.Name)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}
