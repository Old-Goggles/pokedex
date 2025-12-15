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
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        // we’ll add more later
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        // first, check lengths match
        if len(actual) != len(c.expected) {
            t.Errorf("unexpected length for input %q: got %d, want %d",
                c.input, len(actual), len(c.expected))
            continue
        }

        // then, check each word
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("unexpected word at index %d for input %q: got %q, want %q",
                    i, c.input, actual[i], c.expected[i])
            }
        }
    }
}
