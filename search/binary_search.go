package main

import (
	"log"
	"math/rand"
	"time"
)

// Iterative binary search implementation
// returns position of found element
// otherwise -1
func Search(arr []int, el int) int {
	hi := len(arr)
	lo := 0
	for lo < hi {
		mid := (hi + lo) / 2
		if arr[mid] == el {
			return mid
		} else if arr[mid] < el {
			lo = mid
		} else {
			hi = mid
		}
	}
	return -1
}

func main() {
	log.Println("Binary search algorithm demo")
	size := 1000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	rand.Seed(time.Now().Unix())
	element := rand.Intn(size)
	log.Printf("Try find element %d in range %d to %d", element, arr[0], arr[len(arr)-1])
	pos := Search(arr, element)
	if pos != -1 {
		log.Printf("Found element %d by index %d", element, pos)
	} else {
		log.Printf("Element %d not found in array", element)
	}

}
