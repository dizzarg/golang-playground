package chapter1

import (
	"testing"
)

func TestIsRotation(t *testing.T) {
	cases := []struct {
		substr   string
		text     string
		expected bool
	}{
		{"waterbottle", "terbottlewa", true},
		{"hellomynameis", "nameishellomy", true},
		{"waterbottle", "water", false},
		{"waterbottle", "elttobretaw", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := IsRotation(c.substr, c.text)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %t, actual: %t\n",
				c.substr, c.text, c.expected, actual)
		}
	}
}
