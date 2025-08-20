package generate

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	template "github.com/faizisyellow/soho/internal/template/service"
	"github.com/faizisyellow/soho/internal/utils"
	"github.com/pelletier/go-toml/v2"
)

type TomlConfig struct {
	Handler    string
	Router     string
	Service    string
	Repository string
}

const (
	TOMLFILENAME    = ".soho.toml"
	SERVICEFILENAME = "service.go"
)

func GenerateService(name, rootPath string, withTest bool) error {

	var tmlCfg TomlConfig

	tomlf, err := os.ReadFile(filepath.Join(rootPath, TOMLFILENAME))
	if err != nil {
		return err
	}

	err = toml.Unmarshal(tomlf, &tmlCfg)
	if err != nil {
		return err
	}

	serviceTemp := template.ServiceTemplate{
		Name: name,
	}

	dat, err := serviceTemp.NewService(rootPath)
	if err != nil {
		return err
	}

	var filename strings.Builder
	filename.WriteString(strings.ToLower(name))
	filename.WriteString(".go")

	genFile, err := os.Create(filepath.Join(rootPath, tmlCfg.Service, filename.String()))
	if err != nil {
		return err
	}

	_, err = genFile.Write(dat)
	if err != nil {
		return err
	}

	serviceInterfacePath := filepath.Join(rootPath, tmlCfg.Service, SERVICEFILENAME)

	file, err := os.Open(serviceInterfacePath)
	if err != nil {
		return err
	}

	defer file.Close()

	lineNum, err := utils.FindBlockEnd(file, "Service struct")
	if err != nil {
		return err
	}

	datMap, err := serviceTemp.ServiceMapTemplate()
	if err != nil {
		return err
	}

	if err := AppendData(serviceInterfacePath, datMap, lineNum); err != nil {
		return err
	}

	reFile, err := os.Open(serviceInterfacePath)
	if err != nil {
		return err
	}

	defer reFile.Close()

	lineNumServiceImp, err := utils.FindBlockEnd(reFile, "return &Service{")
	if err != nil {
		return err
	}

	datImp, err := serviceTemp.ServiceImplementation()
	if err != nil {
		return err
	}

	if err := AppendData(serviceInterfacePath, datImp, lineNumServiceImp); err != nil {
		return err
	}

	if withTest {
		var filenameTest strings.Builder
		filenameTest.WriteString(strings.ToLower(name))
		filenameTest.WriteString("_test")
		filenameTest.WriteString(".go")

		fileTest, err := os.Create(filepath.Join(rootPath, tmlCfg.Service, filenameTest.String()))
		if err != nil {
			return err
		}

		testTemp, err := serviceTemp.ServiceTestTemplate()
		if err != nil {
			return err
		}

		_, err = fileTest.Write(testTemp)
		if err != nil {
			return err
		}

	}

	return nil
}

func AppendData(path string, data []byte, lineNum int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Validate line index
	if lineNum < 0 || lineNum > len(lines) {
		return fmt.Errorf("line number %d out of bounds", lineNum)
	}

	// Convert inserted block into multiple lines
	newLines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	// Insert at lineNum
	lines = append(lines[:lineNum], append(newLines, lines[lineNum:]...)...)

	// Join back into text
	output := strings.Join(lines, "\n")

	return os.WriteFile(path, []byte(output), 0644)
}
