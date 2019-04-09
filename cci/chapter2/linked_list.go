package chapter2

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *node
	tail *node
}

type node struct {
	data int
	next *node
	prev *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func (list *LinkedList) Append(value int) {
	newNode := &node{value, nil, nil}
	if list.head == nil {
		list.head, list.tail = newNode, newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

func (list *LinkedList) Get(index uint) (int, error) {
	if list.head == nil {
		return 0, errors.New("list nil")
	}
	root := list.head
	if index == 0 {
		return root.data, nil
	}
	var count uint = 0
	for root != list.tail {
		root = root.next
		count++
		if index == count {
			list.head = root.next
			return root.data, nil
		}
	}
	return 0, errors.New("Index out of range")
}

func (list *LinkedList) ForEach(cb func(int)) {
	if list.head == nil {
		return
	}
	root := list.head
	cb(root.data)
	for root != list.tail {
		root = root.next
		cb(root.data)
	}
}

func (list *LinkedList) String() string {
	if list.head == nil {
		return "[]"
	}
	head := fmt.Sprint("[")
	content := ""
	root := list.head
	for root.next != nil {
		content = content + fmt.Sprint(root.data) + ", "
		root = root.next
	}
	content = content + fmt.Sprint(root.data)
	footer := fmt.Sprint("]")
	return head + content + footer
}
