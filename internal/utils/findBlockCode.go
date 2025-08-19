package utils

import (
	"bufio"
	"io"
	"strings"
)

// FindBlockCode finds block of code from textTerm and return
// the close-brace's line of the block and nil if there's no error.
func FindBlockCode(r io.Reader, textTerm string) (int, error) {

	var lineNum = 1
	var endOfCodeLine int
	var braces int
	var inTerm bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, textTerm) {
			inTerm = true
			braces = 0 // reset for new block
		}

		if inTerm {

			braces += strings.Count(line, "{")
			braces -= strings.Count(line, "}")

			// when balanced again, found the end
			if braces == 0 {
				endOfCodeLine = lineNum
				inTerm = false
			}
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return endOfCodeLine, nil
}
