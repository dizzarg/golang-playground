package chapter2

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	if list.head != nil {
		t.Fatalf("Head should be nil. Actual: %v\n", list.head)
	}
	if list.tail != nil {
		t.Fatalf("Tail should be nil. Actual: %v\n", list.tail)
	}
}

func TestAppendElement(t *testing.T) {
	list := NewLinkedList()
	expectedData := 1
	list.Append(expectedData)
	if list.head == nil {
		t.Fatal("Head shouldn't be nil.")
	}
	if list.tail == nil {
		t.Fatal("Tail shouldn't be nil.")
	}
	if list.head != list.tail {
		t.Fatalf("Head and tail should be referred the same node. Head: %v. Tail: %v", list.head, list.tail)
	}
	if list.head.data != expectedData {
		t.Fatalf("Expected: %v, actual: %v\n", expectedData, list.head.data)
	}
}
