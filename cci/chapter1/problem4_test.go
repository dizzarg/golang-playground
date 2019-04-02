package chapter1

import "testing"

func TestCanGeneratePermutations(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},      // normal palindrome
		{"racacer", true},      // jumbled palindrome
		{"abracadabra", false}, // not a palindrome
		{"Tact Coa", true},     // normal palindrome
		{"a", true},
		{"", true},
	}
	for _, c := range cases {
		actual := CanGeneratePermutations(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
