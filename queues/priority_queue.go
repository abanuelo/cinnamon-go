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

// PriorityQueue implements a priority queue based on a min-heap.
type PriorityQueue struct {
	items []*Item
	mu    sync.Mutex
}

// NewPriorityQueue creates a new empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items: make([]*Item, 0),
	}
}

// Len returns the number of items in the priority queue.
func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

// Less returns whether the item with index i has higher priority than the item with index j.
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].Priority < pq.items[j].Priority
}

// Swap swaps the items with indexes i and j in the priority queue.
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	pq.items = append(pq.items, item)
}

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.items)
	item := pq.items[n-1]
	pq.items = pq.items[0 : n-1]
	return item
}

// Enqueue adds an item to the priority queue.
func (pq *PriorityQueue) Enqueue(item *Item) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	heap.Push(pq, item)
}

// Dequeue removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Dequeue() *Item {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*Item)
}

// PrintContents prints the contents of the priority queue
func (pq *PriorityQueue) PrintContents() {
	fmt.Println("Priority Queue Contents:")
	for i := 0; i < pq.Len(); i++ {
		item := pq.items[i]
		fmt.Printf("Index %d: Url: %s, Priority: %d\n", i, item.Url, item.Priority)
	}
}

func TimeoutItems(pq *PriorityQueue, maxAge time.Duration) {
	// TODO REMOVE THE NEED TO DO THIS EVERY 4 SECONDS
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			if len(pq.items) > 0 && currentTime.Sub(((pq.items)[0].Arrival)) > maxAge {
				item := pq.Dequeue()
				item.Status = 1
			}
		}
	}
}
