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
			input:    " Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " My friend is gone. Bye.   ",
			expected: []string{"my", "friend", "is", "gone.", "bye."},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		want := c.expected
		if len(got) != len(want) {
			t.Errorf("the expected and input lenght don't match. Expected: %v, actual: %v", want, got)
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("the words doesn't match.  Expected: %v, actual: %v", want[i], got[i])
			}

		}
	}
}
