package chapter2

import (
	"errors"
	"fmt"
)

// Find kth element from tail of list
// return error if list size less then k and list is empty
func (list *LinkedList) KFromTail(k int) (int, error) {
	if list.head == nil {
		return 0, errors.New("Linked list is empty.")
	}
	current := list.head
	for k > 0 {
		if current == nil {
			return 0, fmt.Errorf("Linked list size less then input value.")
		}
		current = current.next
		k--
	}
	runner := list.head
	for current != nil {
		current = current.next
		runner = runner.next
	}
	return runner.data, nil
}
