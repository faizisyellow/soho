package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	template "github.com/faizisyellow/soho/internal/template/repository"
	"github.com/faizisyellow/soho/internal/utils"
	"github.com/pelletier/go-toml/v2"
)

const (
	REPOSITORYFILENAME = "repository.go"
)

func GenerateRepository(name, root string, isWithTest bool) error {

	var tomlcfg TomlConfig

	tomlf, err := os.ReadFile(filepath.Join(root, TOMLFILENAME))
	if err != nil {
		return err
	}

	err = toml.Unmarshal(tomlf, &tomlcfg)
	if err != nil {
		return err
	}

	repositoryTemp := template.RepositoryTemplate{
		Name: name,
	}

	repoDat, err := repositoryTemp.NewRepository()
	if err != nil {
		return err
	}

	var filename strings.Builder
	filename.WriteString(strings.ToLower(repositoryTemp.Name))
	filename.WriteString(".go")

	genFile, err := os.Create(filepath.Join(root, tomlcfg.Repository, filename.String()))
	if err != nil {
		return err
	}

	defer genFile.Close()

	_, err = genFile.Write(repoDat)
	if err != nil {
		return err
	}

	repositoryInterfacePath := filepath.Join(root, tomlcfg.Repository, REPOSITORYFILENAME)

	file, err := os.Open(repositoryInterfacePath)
	if err != nil {
		return err
	}

	defer file.Close()

	lineNum, err := utils.FindBlockEnd(file, "Repository struct")
	if err != nil {
		return err
	}

	datMap, err := repositoryTemp.RepositoryMap()
	if err != nil {
		return err
	}

	err = AppendData(repositoryInterfacePath, datMap, lineNum)
	if err != nil {
		return err
	}

	reFile, err := os.Open(repositoryInterfacePath)
	if err != nil {
		return err
	}

	defer reFile.Close()

	lineNumRepoImp, err := utils.FindBlockEnd(reFile, "return &Repository{")
	if err != nil {
		return err
	}

	datImp, err := repositoryTemp.RepositoryImplementation()
	if err != nil {
		return err
	}

	if err := AppendData(repositoryInterfacePath, datImp, lineNumRepoImp); err != nil {
		return err
	}

	if isWithTest {
		var filenameTestTemp strings.Builder
		filenameTestTemp.WriteString(strings.ToLower(repositoryTemp.Name))
		filenameTestTemp.WriteString("_test")
		filenameTestTemp.WriteString(".go")

		f, err := os.Create(filepath.Join(root, tomlcfg.Repository, filenameTestTemp.String()))
		if err != nil {
			return err
		}

		defer f.Close()

		dataTest, err := repositoryTemp.RepositoryTest()
		if err != nil {
			return err
		}

		_, err = f.Write(dataTest)
		if err != nil {
			return err
		}

		defer fmt.Printf("repository test created at %v\n", filepath.Join(root, tomlcfg.Repository))
	}

	fmt.Printf("repository created at %v\n", filepath.Join(root, tomlcfg.Repository))

	return nil
}
