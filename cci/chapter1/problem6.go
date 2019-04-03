package chapter1

import (
	"strconv"
)

func BasicCompress(input string) string {
	var compStr string
	counter := 1
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			counter++
		} else {
			compStr = compStr + string(input[i]) + strconv.Itoa(counter)
			counter = 1
		}
	}
	compStr = compStr + string(input[len(input)-1]) + strconv.Itoa(counter)

	if len(compStr) > len(input) {
		return input
	} else {
		return compStr
	}
}
