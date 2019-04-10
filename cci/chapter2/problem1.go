package chapter2

// remove duplicates from linked list
// uses hash map for buffered element
// O(n) complexity
func (list *LinkedList) RemoveDuplicates1() {
	if list.head == nil {
		return
	}
	seen := make(map[int]int)
	root := list.head
	for root != list.tail {
		if _, ok := seen[root.data]; ok {
			list.deleteNode(root)
		}
		seen[root.data]++
		root = root.next
	}
}

// remove duplicates from linked list
// uses two link and two loop
// O(n^2) complexity
func (list *LinkedList) RemoveDuplicates2() {
	if list.head == nil {
		return
	}
	current := list.head
	for current != list.tail {
		runner := current.next
		for runner != list.tail {
			if runner.data == current.data {
				list.deleteNode(runner)
			}
			runner = runner.next
		}
		current = current.next
	}
}
