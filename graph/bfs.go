package graph

import "container/list"

// breadth-first search algorithm
func search(graph map[string]string, name string) {
	searchQueue := list.New()
	searchQueue.PushBack(graph[name])
	searched := make([]string, 16)
	for searchQueue.Len() == 0 {
		person := searchQueue.Back()
	}
}
