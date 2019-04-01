package chapter1

// The function checks permutation.
// if one is a permutation of the other returns true otherwise false.
func ArePermutations(text1, text2 string) bool {
	if len(text1) != len(text2) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range text1 {
		counts[r]++
	}
	for _, r := range text2 {
		counts[r]--
	}
	for _, val := range counts {
		if val != 0 {
			return false
		}
	}
	return true
}
