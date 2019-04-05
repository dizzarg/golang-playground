package chapter1

import (
	"strings"
)

// checked if substr contains in a text
func IsRotation(substr, text string) bool {
	if len(substr) != len(text) {
		return false
	}
	newStr := text + text
	return strings.Contains(newStr, substr)
}
