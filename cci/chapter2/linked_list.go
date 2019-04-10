package chapter2

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	data int
	prev *Node
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{value, nil, nil}
	if list.head == nil {
		list.head, list.tail = newNode, newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

func (list *LinkedList) Delete(index uint) error {
	if node, err := list.GetNode(index); err == nil {
		list.deleteNode(node)
		node = nil
		return nil
	} else {
		return err
	}
}

func (list *LinkedList) deleteNode(node *Node) {
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	} else {
		list.head = next
	}
	if next != nil {
		next.prev = prev
	} else {
		list.tail = prev
	}
	node = nil
}

func (list LinkedList) GetNode(index uint) (*Node, error) {
	if list.head == nil {
		return nil, errors.New("list nil")
	}
	root := list.head
	if index == 0 {
		return root, nil
	}
	var count uint = 0
	for root != list.tail {
		root = root.next
		count++
		if index == count {
			list.head = root.next
			return root, nil
		}
	}
	return nil, errors.New("Index out of range")
}

func (list LinkedList) GetValue(index uint) (int, error) {
	if node, err := list.GetNode(index); err == nil {
		return node.data, nil
	} else {
		return -1, err
	}
}

func (list LinkedList) ForEach(cb func(uint, *Node)) {
	if list.head == nil {
		return
	}
	var index uint = 0
	root := list.head
	cb(index, root)
	for root != list.tail {
		root = root.next
		index++
		cb(index, root)
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

func (node *Node) String() string {
	return fmt.Sprintf("%p -> (%d, %p)-> %p", node.prev, node.data, node, node.next)
}
