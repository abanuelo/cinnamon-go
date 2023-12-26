package inflight

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/abanuelo/cinnamon-go/priorityq"
)

// Inflight Request limit
var INFLIGHT_LIMIT = 10
var inflight = 0

func processRequest(req priorityq.Item) {
	// Simulate processing time
	time.Sleep(2 * time.Second)
	fmt.Printf("HERE: Processed request %d: %s\n", req.Priority, req.Value)
	inflight -= 1
}

func Worker(pq *priorityq.PriorityQueue, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Pop the highest-priority request from the priority queue
		if pq.Len() > 0 && inflight+1 < INFLIGHT_LIMIT {
			inflight += 1
			fmt.Printf("Current inflight: %d\n", inflight)
			req := heap.Pop(pq).(*priorityq.Item)
			go processRequest(*req)
		} else {
			fmt.Printf("Inflight limit reached or priority queue empty: %d\n", inflight)
		}
		// Add a sleep or some mechanism to avoid busy-waiting
		time.Sleep(1 * time.Second)
	}
}
