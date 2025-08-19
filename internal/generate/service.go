package generate

import (
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
	FILENAME        = ".soho.toml"
	SERVICEFILENAME = "service.go"
)

func GenerateService(name, rootPath string, withTest bool) error {

	var tmlCfg TomlConfig
	tomlf := make([]byte, 90)

	f, err := os.Open(filepath.Join(rootPath, FILENAME))
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Read(tomlf)
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
	filename.WriteString(".go.txt")

	genFile, err := os.Create(filepath.Join(rootPath, tmlCfg.Service, filename.String()))
	if err != nil {
		return err
	}

	_, err = genFile.Write(dat)
	if err != nil {
		return err
	}

	// Adding new service map
	serviceDir, err := os.ReadDir(filepath.Join(rootPath, tmlCfg.Service))
	if err != nil {
		return err
	}

	var isServiceDirFound bool

	for _, serv := range serviceDir {

		if serv.Name() == SERVICEFILENAME {
			isServiceDirFound = true
		}
	}

	if !isServiceDirFound {
		return fmt.Errorf("service file not found")
	}

	serviceInterfaceFile, err := os.OpenFile(filepath.Join(rootPath, tmlCfg.Service, SERVICEFILENAME), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer serviceInterfaceFile.Close()

	lastBlockServiceMap, err := utils.FindBlockCode(serviceInterfaceFile, "interface")
	if err != nil {
		return err
	}

	datMap, err := serviceTemp.ServiceMapTemplate()
	if err != nil {
		return err
	}

	// _, err = serviceInterfaceFile.WriteAt(datMap, int64(lastBlockServiceMap+1))
	// if err != nil {
	// 	return err
	// }

	// TODO: append data to existing file
	err = AppendData(filepath.Join(rootPath, tmlCfg.Service, SERVICEFILENAME), datMap, lastBlockServiceMap+1)
	if err != nil {
		return err
	}

	// lastBlockServiceImp, err := utils.FindBlockCode(serviceInterfaceFile, "return")
	// if err != nil {
	// 	return err
	// }

	// datImp, err := serviceTemp.ServiceImplementation()
	// if err != nil {
	// 	return err
	// }

	// _, err = serviceInterfaceFile.WriteAt(datImp, int64(lastBlockServiceImp))
	// if err != nil {
	// 	return err
	// }

	if withTest {
		var filenameTest strings.Builder
		filenameTest.WriteString(name)
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

func AppendData(path string, data []byte, at int) error {

	// Read all file content
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Prevent out-of-bounds index
	if at < 0 || at > len(content) {
		return fmt.Errorf("index %d out of bounds", at)
	}

	// Insert the new data at position `at`
	newContent := append(content[:at], append(data, content[at:]...)...)

	// Write back to the file (overwrite)
	return os.WriteFile(path, newContent, 0644)
}
