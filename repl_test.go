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
			input:    "         ",
			expected: []string{},
		},
		{
			input:    "   Hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLOWORLD",
			expected: []string{"helloworld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

        if len(actual) != len(c.expected) {
            t.Errorf("Actual length: %d\t|\tExpected length: %d", len(actual), len(c.expected))
            t.Fail()
        }

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
                t.Errorf("Got: %s\t|\tExpecting: %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
