package main

import "testing"


func TestCleanText(t *testing.T) {
    cases := []struct{
        input    string
        expected []string
    }{
        {
            input: "  hello   world  ",
            expected: []string{"hello", "world"},
        },
        {
			input:    "  foo bar   baz  ",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "singleword",
			expected: []string{"singleword"},
		},
		{
			input:    "     ",
			expected: []string{},
		},
		{
			input:    "  a  b c   ",
			expected: []string{"a", "b", "c"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
            continue
        }

        for i:= range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("For input '%s', expected '%s' at index %d but got '%s'", c.input, c.expected[i], i, actual[i])
            }
        }
    }
}

