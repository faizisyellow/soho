package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// FindBlockEnd finds the line number just before the closing brace
// of the block that starts at the given textTerm.
// Example: if term is "interface" -> returns last line inside interface block.
//
//	if term is "return &Service{" -> returns last field line inside the return block.
func FindBlockEnd(r io.Reader, textTerm string) (int, error) {
	var lineNum = 1
	var endLine = -1
	var braces int
	var inBlock bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, textTerm) && !inBlock {
			inBlock = true
			braces = strings.Count(line, "{") - strings.Count(line, "}")
		} else if inBlock {
			braces += strings.Count(line, "{")
			braces -= strings.Count(line, "}")

			if braces == 0 {
				endLine = lineNum
				break
			}
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}
	if !inBlock {
		return -1, fmt.Errorf("term %q not found", textTerm)
	}
	if endLine == -1 {
		return -1, fmt.Errorf("block for term %q not closed", textTerm)
	}

	// return the line *before* the closing brace
	return endLine - 1, nil
}
