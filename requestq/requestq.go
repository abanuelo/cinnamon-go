package requestq

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
	time.Sleep(10 * time.Second)
	fmt.Printf("HERE: Processed request %d: %s\n", req.Priority, req.Value)
	inflight -= 1
}

func Worker(worker int, pq *priorityq.PriorityQueue, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		if inflight+1 < INFLIGHT_LIMIT {
			mu.Lock()
			if pq.Len() > 0 {
				inflight += 1
				fmt.Printf("Current inflight: %d with worker id: %d\n", inflight, worker)
				req := heap.Pop(pq).(*priorityq.Item)
				req.Processed = "processed"
				mu.Unlock()
				go processRequest(*req)
			} else {
				mu.Unlock()
			}
		}
	}
}
