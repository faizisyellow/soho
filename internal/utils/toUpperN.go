package utils

import (
	"strings"
	"unicode"
)

func ToUpperN(text string, n int) string {

	var newForm strings.Builder

	for i, char := range text {
		if i == n {
			if !unicode.IsUpper(char) {
				newForm.WriteRune(unicode.ToUpper(char))

				continue
			}
		}

		newForm.WriteRune(char)
	}

	return newForm.String()
}
