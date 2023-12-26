package priorityq

import "fmt"

// Item represents an item in the priority queue.
type Item struct {
	Value    interface{}
	Priority int
}

// PriorityQueue is a min-heap implementation.
type PriorityQueue []*Item

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items based on priority.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps two items in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop removes and returns the minimum item from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) PrintContents() string {
	var result string

	for _, item := range pq {
		result += fmt.Sprintf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}

	return result
}
