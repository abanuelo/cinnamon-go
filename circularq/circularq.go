package circularq

import (
	"fmt"
	"sort"
	"sync"
)

// CircularQueue represents a circular queue structure of integers
type CircularQueue struct {
	items []int
	size  int
	head  int
	tail  int
	mu    sync.Mutex
}

// NewCircularQueue creates a new circular queue with the specified size
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, size),
		size:  size,
		head:  0,
		tail:  -1,
	}
}

// CurrentCapacity returns the current capacity of the circular queue
func (cq *CircularQueue) CurrentCapacity() int {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	// Calculate the number of elements in the queue
	count := 0
	for i := 0; i < cq.size; i++ {
		idx := (cq.head + i) % cq.size
		if cq.items[idx] != 0 {
			count++
		}
	}

	return count
}

// Enqueue adds an item to the queue, bumping out the oldest item if necessary
func (cq *CircularQueue) Enqueue(item int) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.tail = (cq.tail + 1) % cq.size
	cq.items[cq.tail] = item
}

// PrintQueue prints the items in the queue
func (cq *CircularQueue) PrintQueue() {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	fmt.Println("Queue items:")
	for i := 0; i < cq.size; i++ {
		idx := (cq.head + i) % cq.size
		fmt.Println(cq.items[idx])
	}
}

// PercentileDistribution calculates the percentile distribution based on the items in the queue
func (cq *CircularQueue) PercentileDistribution(percentile float64) int {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	// Create a sorted copy of the items
	sortedItems := make([]int, len(cq.items))
	copy(sortedItems, cq.items)
	sort.Ints(sortedItems)

	// Calculate the values at the specified percentiles
	index := int(float64(len(sortedItems)-1) * percentile / 100.0)
	value := sortedItems[index]

	return value
}
