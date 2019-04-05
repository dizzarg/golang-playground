package chapter1

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	cases := []struct {
		testMatrix [][]int
		expected   [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			[][]int{
				{1, 2, 0},
				{4, 5, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{0, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{0, 2, 3},
				{0, 0, 0},
				{0, 8, 9},
			},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}
	for _, c := range cases {
		SetZeros(c.testMatrix)
		if !reflect.DeepEqual(c.testMatrix, c.expected) {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, c.testMatrix)
		}
	}
}
