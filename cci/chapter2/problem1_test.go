package chapter2

import "testing"

func TestLinkedList_RemoveDuplicates1_OrderedList(t *testing.T) {
	vals := []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	list := NewLinkedList()
	for _, v := range vals {
		list.Append(v)
	}
	prev := 0
	list.RemoveDuplicates1()
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
		if prev > node.data {
			t.Fatalf("Prev value must less than next. Prev: %v, next: %v\n", prev, node.data)
		}
	})
	if count != 6 {
		t.Fatalf("List must have %d elements. Actual: %d\n", 6, count)
	}
}

func TestLinkedList_RemoveDuplicates1_UnorderedList(t *testing.T) {
	vals := []int{1, 1, 3, 4, 5, 1, 2, 3, 5, 6}
	list := NewLinkedList()
	for _, v := range vals {
		list.Append(v)
	}
	prev := 0
	list.RemoveDuplicates1()
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
		if prev > node.data {
			t.Fatalf("Prev value must less than next. Prev: %v, next: %v\n", prev, node.data)
		}
	})
	if count != 6 {
		t.Fatalf("List must have %d elements. Actual: %d\n", 6, count)
	}
}

func TestLinkedList_RemoveDuplicates2_UnorderedList(t *testing.T) {
	vals := []int{1, 1, 3, 4, 5, 1, 2, 3, 5, 6}
	list := NewLinkedList()
	for _, v := range vals {
		list.Append(v)
	}
	prev := 0
	list.RemoveDuplicates2()
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
		if prev > node.data {
			t.Fatalf("Prev value must less than next. Prev: %v, next: %v\n", prev, node.data)
		}
	})
	if count != 6 {
		t.Fatalf("List must have %d elements. Actual: %d\n", 6, count)
	}
}
