package chapter2

// remove duplicates from linked list
func (list *LinkedList) RemoveDuplicates() {
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
