package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	template "github.com/faizisyellow/soho/internal/template/handler"
	"github.com/pelletier/go-toml/v2"
)

const (
	HANDLERFILENAME = "handler.go"
)

func GenerateHandler(name, root string, isWithTest bool) error {

	var tomlcfg TomlConfig

	tomlf, err := os.ReadFile(filepath.Join(root, TOMLFILENAME))
	if err != nil {
		return err
	}

	err = toml.Unmarshal(tomlf, &tomlcfg)
	if err != nil {
		return err
	}

	handlerTemp := template.HandlerTemplate{
		Name: name,
	}

	handlerDat, err := handlerTemp.Handler()
	if err != nil {
		return err
	}

	var filename strings.Builder
	filename.WriteString(strings.ToLower(handlerTemp.Name))
	filename.WriteString(".go")

	genFile, err := os.Create(filepath.Join(root, tomlcfg.Handler, filename.String()))
	if err != nil {
		return err
	}

	_, err = genFile.Write(handlerDat)
	if err != nil {
		return err
	}

	if isWithTest {
		var filenameTestTemp strings.Builder
		filenameTestTemp.WriteString(strings.ToLower(handlerTemp.Name))
		filenameTestTemp.WriteString("_test")
		filenameTestTemp.WriteString(".go")

		f, err := os.Create(filepath.Join(root, tomlcfg.Handler, filenameTestTemp.String()))
		if err != nil {
			return err
		}

		defer f.Close()

		dataTest, err := handlerTemp.HandlerTest()
		if err != nil {
			return err
		}

		_, err = f.Write(dataTest)
		if err != nil {
			return err
		}

	}

	fmt.Printf("created handler at %v", filepath.Join(root, tomlcfg.Handler))

	if isWithTest {
		fmt.Printf("created handler test at %v", filepath.Join(root, tomlcfg.Handler))
	}

	return nil
}
