package chapter2

import "testing"

func TestLinkedList_KFromTail(t *testing.T) {
	list := NewList([]int{10, 1, 3, 4, 5, 1, 2, 3, 5, 6})
	cases := []struct {
		pos    int
		expect int
	}{
		{5, 1},
		{1, 6},
		{10, 10},
	}

	for _, c := range cases {
		val, err := list.KFromTail(c.pos)
		if err != nil {
			t.Fatalf("Function must find %d-th element from tail", c.pos)
		}
		if val != c.expect {
			t.Fatalf("%d-th mast expect %d, but actual %d", c.pos, c.expect, val)
		}
	}
}

func TestLinkedList_KFromTail_emptyList(t *testing.T) {
	list := EmptyList()
	pos := 1
	val, err := list.KFromTail(pos)
	if err == nil {
		t.Fatalf("Function must find %d-th element from tail", pos)
	}
	if val != 0 {
		t.Fatalf("%d-th mast expect 0, but actual %d", pos, val)
	}
}
