package queues

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Needed to define to remove circular dependencies
type Item struct {
	Method   string
	Url      string
	Priority int64
	Arrival  time.Time
	Status   int
}

// PriorityQueue represents a min-heap priority queue.
type PriorityQueue []*Item

// Len returns the number of items in the priority queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less defines the priority order of items in the priority queue.
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

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// NewPriorityQueue creates a new empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq)
	return &pq
}

// RemoveOldItems removes items older than maxAge from the priority queue.
var mtx sync.Mutex

func TimeoutItems(pq *PriorityQueue, maxAge time.Duration) {
	for {
		select {
		case <-time.After(maxAge):
			currentTime := time.Now()

			mtx.Lock()
			fmt.Println("**********************************")
			fmt.Println("[START] - TimeoutItems")
			for len(*pq) > 0 && currentTime.Sub(((*pq)[0].Arrival)) > maxAge {
				item := heap.Pop(pq).(*Item)
				fmt.Println("Popped Item")
				item.Status = 1
			}
			fmt.Println("[END] - TimeoutItems")
			fmt.Println("**********************************")
			mtx.Unlock()
		}
	}
}
