package main

// Iterative binary search implementation
// returns position of found element
// otherwise -1
func BinarySearch(arr []int, el int) int {
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

// Iterative linear search implementation
// returns position of found element
// otherwise -1
func LinearSearch(arr []int, el int) int {
	for i, v := range arr {
		if v == el {
			return i
		}
	}
	return -1
}
