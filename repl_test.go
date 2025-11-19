package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "                  test",
			expected: []string{"test"},
		},
		{
			input:    "          t e s t  ",
			expected: []string{"t", "e", "s", "t"},
		},
		{
			input:    "             ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("lengths don't match: '%v' '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words don't match: '%v' '%v'", word, expectedWord)
				continue
			}
		}
	}
}
