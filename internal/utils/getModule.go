package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func GetModuleName(root string) (string, error) {

	file, err := os.Open(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}

	defer file.Close()

	f := bufio.NewReader(file)

	b, _, err := f.ReadLine()
	if err != nil {
		return "", err
	}
	return string(b[7:]), nil
}
