package chapter1

// The function replace all spaces in a string with '%20'.
func URLify(url string) string {
	newStringLen := 0
	for _, v := range url {
		newStringLen++
		if v == ' ' {
			// +2 for each space for the extra "20"s.
			newStringLen += 2
		}
	}
	result := make([]rune, newStringLen)
	i := 0
	for _, v := range url {
		if v == ' ' {
			result[i] = rune('%')
			result[i+1] = rune('2')
			result[i+2] = rune('0')
			i += 3
		} else {
			result[i] = v
			i++
		}
	}
	return string(result)
}
