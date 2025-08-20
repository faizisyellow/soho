package utils

import "testing"

func TestUpperN(t *testing.T) {

	testsCase := []struct {
		name    string
		text    string
		n       int
		expectd string
	}{
		{
			name:    "should upper first character",
			text:    "lizzy",
			n:       0,
			expectd: "Lizzy",
		},
		{
			name:    "it is already upper",
			text:    "Dakota",
			n:       0,
			expectd: "Dakota",
		},
		{
			name:    "should upper char in the middle",
			text:    "monroes",
			n:       3,
			expectd: "monRoes",
		},
	}

	for _, tc := range testsCase {
		t.Run(tc.name, func(t *testing.T) {

			res := ToUpperN(tc.text, tc.n)
			if res != tc.expectd {
				t.Errorf("expected %v but got %v", tc.expectd, res)
			}
		})
	}
}
