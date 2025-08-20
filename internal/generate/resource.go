package generate

import (
	"os"
	"path/filepath"

	template "github.com/faizisyellow/soho/internal/template/route"
	"github.com/faizisyellow/soho/internal/utils"
	"github.com/pelletier/go-toml/v2"
)

const ROUTEFILENAME = "mux.go"

func GenerateResource(name, root string, isWithTest bool) error {

	err := GenerateRepository(name, root, isWithTest)
	if err != nil {
		return err
	}

	err = GenerateService(name, root, isWithTest)
	if err != nil {
		return err
	}

	err = GenerateHandler(name, root, isWithTest)
	if err != nil {
		return err
	}

	var tmlCfg TomlConfig

	tomlf, err := os.ReadFile(filepath.Join(root, TOMLFILENAME))
	if err != nil {
		return err
	}

	err = toml.Unmarshal(tomlf, &tmlCfg)
	if err != nil {
		return err
	}

	routeTemp := template.RouteTemplate{
		Name: name,
	}

	routeDat, err := routeTemp.NewRoute()
	if err != nil {
		return err
	}

	routerPath := filepath.Join(root, tmlCfg.Router, ROUTEFILENAME)

	file, err := os.Open(routerPath)
	if err != nil {
		return err
	}

	defer file.Close()

	lineNum, err := utils.FindBlockEnd(file, `r.Route("/v1", func(r chi.Router) {`)
	if err != nil {
		return err
	}

	if err := AppendData(routerPath, routeDat, lineNum); err != nil {
		return err
	}

	return nil
}
