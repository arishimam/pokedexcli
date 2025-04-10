package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hELLoworld ",
			expected: []string{"helloworld"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "   leading and middle   spaces ",
			expected: []string{"leading", "and", "middle", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of actual: %v, length of expected: %v, not equal", len(actual), len(c.expected))
			t.Fail()
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual: %v and expected: %v are not equal!", word, expectedWord)
				t.Fail()
			}

		}

	}
}
