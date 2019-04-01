package chapter1

import (
	"testing"
)

func TestArePermutations(t *testing.T) {
	cases := []struct {
		text1    string
		text2    string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := ArePermutations(c.text1, c.text2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %v, actual: %v\n",
				c.text1, c.text2, c.expected, actual)
		}
	}
}
