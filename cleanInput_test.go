package main

import (
	"testing"

	"github.com/toine08/pokedexcli/utils"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " I am toto",
			expected: []string{"i", "am", "toto"}, // note case sensitivity
		},
		{
			input:    " I am toto and toine",
			expected: []string{"i", "am", "toto", "and", "toine"}, // note case sensitivity
		},
	}
	for _, c := range cases {
		actual := utils.CleanInput(c.input) // note the capital letter for exported function
		if len(actual) != len(c.expected) {
			t.Errorf("Length of slices are not matching")
			continue
		}
		for i, word := range actual {
			if word != c.expected[i] {
				t.Errorf("Mismatch at index %d: got %s, want %s", i, word, c.expected[i])
			}
		}
	}
}
