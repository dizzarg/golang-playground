package main

import "testing"

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
