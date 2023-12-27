package priorityq

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Item represents an item in the priority queue.
type Item struct {
	Value    interface{}
	Priority int
	Arrival  time.Time
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

// Prints queue contents
func (pq PriorityQueue) PrintContents() string {
	var result string

	for _, item := range pq {
		result += fmt.Sprintf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}

	return result
}

func RemoveOldItems(pq *PriorityQueue, maxAge time.Duration, mu *sync.Mutex) {
	for {
		select {
		case <-time.After(maxAge):
			currentTime := time.Now()

			mu.Lock()
			// for len(*pq) > 0 && currentTime.Sub((*pq)[0].Arrival) > maxAge {
			if len(*pq) > 0 && currentTime.Sub((*pq)[0].Arrival) > maxAge {
				item := heap.Pop(pq).(*Item)
				fmt.Printf("TIMEOUT: Removed item: %s, Priority: %d, Arrival: %v\n", item.Value, item.Priority, item.Arrival)
			}
			mu.Unlock()
		}
	}
}
