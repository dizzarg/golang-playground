package chapter2

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	if list.head != nil {
		t.Fatalf("Head must be nil. Actual: %v\n", list.head)
	}
	if list.tail != nil {
		t.Fatalf("Tail must be nil. Actual: %v\n", list.tail)
	}
}

func TestLinkedList_Append(t *testing.T) {
	list := NewLinkedList()
	expectedData := 1
	list.Append(expectedData)
	if list.head == nil {
		t.Fatal("Head mustn't be nil.")
	}
	if list.tail == nil {
		t.Fatal("Tail shouldn't be nil.")
	}
	if list.head != list.tail {
		t.Fatalf("Head and tail must be referred the same Node. Head: %v. Tail: %v", list.head, list.tail)
	}
	if list.head.data != expectedData {
		t.Fatalf("Expected: %v, actual: %v\n", expectedData, list.head.data)
	}
}

func TestLinkedList_ForEach(t *testing.T) {
	list := NewLinkedList()
	prev := 0
	expectElements := 10
	for i := prev + 1; i <= prev+expectElements; i++ {
		list.Append(i)
	}
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
		if prev > node.data {
			t.Fatalf("Prev value must less than next. Prev: %v, next: %v\n", prev, node.data)
		}
	})
	if count != expectElements {
		t.Fatalf("List must have %d elements. Actual: %d\n", expectElements, count)
	}
}

func TestLinkedList_DeleteNode_RemoveCenter(t *testing.T) {

	list := NewLinkedList()
	prev := 0
	expectElements := 10
	for i := prev + 1; i <= prev+expectElements; i++ {
		list.Append(i)
	}
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
	})
	if err := list.Delete(1); err != nil {
		t.Error(err)
	}
	list.ForEach(func(index uint, node *Node) {
		count--

	})
	if count != 1 {
		t.Fatalf("List must have %d elements. Actual: %d\n", expectElements-1, expectElements-count)
	}
}

func TestLinkedList_DeleteNode_RemoveFirst(t *testing.T) {

	list := NewLinkedList()
	prev := 0
	expectElements := 10
	for i := prev + 1; i <= prev+expectElements; i++ {
		list.Append(i)
	}
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
	})
	if err := list.Delete(0); err != nil {
		t.Error(err)
	}
	list.ForEach(func(index uint, node *Node) {
		count--

	})
	if count != 1 {
		t.Fatalf("List must have %d elements. Actual: %d\n %v", expectElements-1, expectElements-count, list)
	}
}

func TestLinkedList_DeleteNode_RemoveLast(t *testing.T) {

	list := NewLinkedList()
	prev := 0
	expectElements := 10
	for i := prev + 1; i <= prev+expectElements; i++ {
		list.Append(i)
	}
	count := 0
	list.ForEach(func(index uint, node *Node) {
		count++
	})
	if err := list.Delete(9); err != nil {
		t.Error(err)
	}
	list.ForEach(func(index uint, node *Node) {
		count--

	})
	if count != 1 {
		t.Fatalf("List must have %d elements. Actual: %d\n %v", expectElements-1, expectElements-count, list)
	}
}
