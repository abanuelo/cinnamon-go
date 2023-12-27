package requestq

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/abanuelo/cinnamon-go/priorityq"
)

// Inflight Request limit
var INFLIGHT_LIMIT float64 = 10.0
var CURR_INFLIGHT float64 = 0.0

func processRequest(req priorityq.Item) {
	// Simulate processing time
	time.Sleep(10 * time.Second)
	fmt.Printf("HERE: Processed request %d: %s\n", req.Priority, req.Value)
	CURR_INFLIGHT -= 1
}

func Worker(worker int, pq *priorityq.PriorityQueue, wg *sync.WaitGroup, mu *sync.Mutex, OUT *float64) {
	defer wg.Done()
	for {
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			mu.Lock()
			if pq.Len() > 0 {
				CURR_INFLIGHT += 1
				fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
				req := heap.Pop(pq).(*priorityq.Item)
				*OUT += 1
				req.Processed = "processed"
				mu.Unlock()
				go processRequest(*req)
			} else {
				mu.Unlock()
			}
		}
		// time.Sleep(2 * time.Second)
	}
}
