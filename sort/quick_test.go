package main

import "testing"

func isSorted(arr []int) bool {
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > prev {
			return false
		}
	}
	return true
}

func TestQuickSort(t *testing.T) {
	data := make([]int, 1<<10)
	for i := 0; i < len(data); i++ {
		data[i] = i ^ 0x2cc
	}
	originData := make([]int, 1<<10)
	copy(originData, data)
	QSort(data)
	if isSorted(data) {
		t.Errorf("sorted %v", originData)
		t.Errorf("   got %v", data)
	}
}

func BenchmarkSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		QSort(data)
		b.StopTimer()
	}
}
