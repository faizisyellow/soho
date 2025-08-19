package generate

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"
)

func TestGenerateNewService(t *testing.T) {

	contentExpected := `package service

			import (
					"context"
					github.com/foobar/"internal/repository"
			)

			type FooServices struct {
				Repository repository.Repository
			}


			func (Foo *FooServices) Create(ctx context.Context) (string, error) {
			
				return "", nil
			}

			func (Foo *FooServices) FindAll(ctx context.Context) ([]repository.FooModel, error) {

				return nil, nil
			}

			func(Foo *FooServices) FindById(ctx context.Context, id int) (repository.FooModel, error) {

				return repository.FooModel{},nil
			}

			func (Foo *FooServices) Update(ctx context.Context, id int, nw repository.FooModel) error {

				return nil
			}

			func (Foo *FooServices) Delete(ctx context.Context, id int) error {

				return nil
			}
	`

	setupTestService(t)

	err := GenerateService("Foo", "", false)
	if err != nil {
		t.Error(err)
		return
	}

	resFile, err := os.Open("Foo.go")
	if err != nil {
		t.Error(err)
		return
	}

	expectedReader := strings.NewReader(contentExpected)

	expectedScan := bufio.NewScanner(expectedReader)
	resultScan := bufio.NewScanner(resFile)

	for resultScan.Scan() {

		expectedScan.Scan()

		if !bytes.Equal(resultScan.Bytes(), expectedScan.Bytes()) {
			t.Errorf("expected same but missmatched")
			return
		}
	}
}

func setupTestService(t *testing.T) {
	t.Helper()

	tomldat := TomlConfig{
		Service: "/service",
	}

	tomlFile, err := os.Create(".soho.toml")
	if err != nil {
		t.Error(err)
	}

	defer tomlFile.Close()

	dat, err := toml.Marshal(tomldat)
	if err != nil {
		t.Error(err)
	}

	_, err = tomlFile.Write(dat)
	if err != nil {
		t.Error(err)
	}

	gomodFile, err := os.Create("go.mod")
	if err != nil {
		t.Error(err)
	}

	defer gomodFile.Close()

	gomodDat := `module github.com/foobar`
	_, err = gomodFile.Write([]byte(gomodDat))
	if err != nil {
		t.Error(err)
	}

	_, err = os.MkdirTemp("", "service")
	if err != nil {
		t.Error(err)
	}

}
