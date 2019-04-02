package chapter1

import "unicode"

// function to check if it is a permutation of a palindrome
// function ignore spaces
// Examples:
// Anna  			-> true
// Civic 			-> true
// Step on no pets	-> true
// Top spot 		-> true
// abracadabra 		-> false
func CanGeneratePermutations(sentence string) bool {
	counts := make(map[rune]int)
	inputLen := 0
	for _, r := range sentence {
		if r == rune(' ') {
			continue
		}
		counts[unicode.ToLower(r)]++
		inputLen++
	}
	odd := inputLen%2 == 1
	seenOdd := false
	for _, val := range counts {
		if val%2 == 1 {
			if !seenOdd && odd {
				seenOdd = true
			} else {
				return false
			}
		}
	}
	return true
}
