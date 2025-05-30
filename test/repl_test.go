package pokedexcli_test

import (
	"testing"

	"github.com/Amarothia/pokedexcli/funcs"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello 		 world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hellO world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hElLo wOrLd",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := funcs.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length: %d != expected length: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual word: %s != expected word: %s", word, expectedWord)
			}
		}
	}
}
