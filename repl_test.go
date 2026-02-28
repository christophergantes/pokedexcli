package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "ALL CAPS",
			expected: []string{"all", "caps"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("\nExpected: %v\nActual: %v", c.expected, actual)
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("\nExpected: %v\nActual: %v", c.expected, actual)
			}
		}
	}
}
