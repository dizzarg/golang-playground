package main

// Iterative binary search implementation
// returns position of found element
// otherwise -1
func Search(arr []int, el int) int {
	hi := len(arr) - 1
	lo := 0
	for lo <= hi {
		mid := (hi + lo) / 2
		if arr[mid] == el {
			return mid
		} else if arr[mid] < el {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
