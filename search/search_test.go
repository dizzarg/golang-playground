package main

import (
	"math/rand"
	"testing"
)

var testArray = [9]int{10, 20, 30, 40, 50, 60, 70, 80, 90}

func TestBinarySearch_successFound(t *testing.T) {
	expectedIndex := 3
	actualIndex := BinarySearch(testArray[:], 40)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func TestBinarySearch_successNotFound(t *testing.T) {
	expectedIndex := -1
	actualIndex := BinarySearch(testArray[:], 200)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func TestBinarySearch_nothingFoundInEmptyArray(t *testing.T) {
	testArray := [0]int{}
	expectedIndex := -1
	actualIndex := BinarySearch(testArray[:], 0)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func TestLinearSearch_successFound(t *testing.T) {
	expectedIndex := 3
	actualIndex := LinearSearch(testArray[:], 40)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func TestLinearSearch_successNotFound(t *testing.T) {
	expectedIndex := -1
	actualIndex := LinearSearch(testArray[:], 200)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func TestLinearSearch_nothingFoundInEmptyArray(t *testing.T) {
	testArray := [0]int{}
	expectedIndex := -1
	actualIndex := LinearSearch(testArray[:], 0)
	if actualIndex != expectedIndex {
		t.Errorf("For %v expected %d, but got %d", testArray, expectedIndex, actualIndex)
	}
}

func BenchmarkBinarySearch1KFindRandom(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i
		}
		el := data[rand.Int31n(1<<10)]
		b.StartTimer()
		BinarySearch(data, el)
		b.StopTimer()
	}
}

func BenchmarkLinearSearch1KFindRandom(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i
		}
		el := data[rand.Int31n(1<<10)]
		b.StartTimer()
		LinearSearch(data, el)
		b.StopTimer()
	}
}
