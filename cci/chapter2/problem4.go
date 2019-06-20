package chapter2

// list petitioner by value function
func (list *LinkedList) PivotAroundValue(value int) {
	lower := EmptyList()
	higher := EmptyList()
	current := list.head
	for current != nil {
		next := current.next
		if current.data <= value {
			lower.Append(current.data)
		} else {
			higher.Append(current.data)
		}
		current = next
	}
	// Join higher and lower lists, and replace.
	lower.tail.next = higher.head
	higher.head.prev = lower.tail
	list.tail = higher.tail
	list.head = lower.head
}
