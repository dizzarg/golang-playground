package main

func isSorted(arr []int) bool {
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < prev {
			return false
		}
	}
	return true
}
