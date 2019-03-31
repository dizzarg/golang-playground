package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sort array of integer values
func QSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

// the recursion function which implement quick sort
// part of array
func sort(arr []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	p := partition(arr, lo, hi)
	sort(arr, lo, p-1)
	sort(arr, p+1, hi)
}

// The function takes last element as pivot, places
// the pivot element at its correct position in sorted
// array, and places all smaller than pivot to left of pivot
// and all greater elements to right of pivot
func partition(arr []int, lo int, hi int) int {
	i := lo + 1
	j := hi
	pivot := arr[lo]
	for {
		for arr[i] < pivot {
			if i == hi {
				break
			}
			i++
		}
		for pivot < arr[j] {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}

func main() {
	fmt.Println("Quick sort algorithm demo")
	size := 20
	arr := make([]int, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Printf("Sorting array \n%v ... \n", arr)
	QSort(arr)
	fmt.Printf("Sorted array \n%v\n", arr)
}
