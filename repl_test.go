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
			input:    " ChArizard porky PIG",
			expected: []string{"charizard", "porky", "pig"},
		},
		// {
		// 	input:    "  lots    of spaces  here ",
		// 	expected: []string{"lots", "of", "spaces", "here"},
		// },
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length of output array is not correct")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match")
			}
		}
	}
}
