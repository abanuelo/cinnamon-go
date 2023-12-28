package cinnamon

import (
	"container/heap"
	"fmt"
	sync "sync"
	"time"

	"github.com/abanuelo/cinnamon-go/queues"
)

func processRequest(req queues.Item) {
	// Simulate processing time
	time.Sleep(5 * time.Second)
	fmt.Printf("HERE: Processed request %d: %s\n", req.Priority, req.Url)
	CURR_INFLIGHT -= 1
}

func Worker(worker int, pq *queues.PriorityQueue, cq *queues.CircularQueue, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			mu.Lock()
			if pq.Len() > 0 {
				// fmt.Printf("WORKING!!!!!!!!!!")
				CURR_INFLIGHT += 1
				fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
				req := heap.Pop(pq).(*queues.Item)
				// cq.Enqueue(int(req.Priority))
				// CURR_INFLIGHT -= 1
				OUT += 1
				req.Status = int(Status_OK)
				mu.Unlock()
				go processRequest(*req)
			} else {
				// fmt.Printf("PQ is empty!!!!!!!!!!")
				mu.Unlock()
			}
		}
		// time.Sleep(2 * time.Second)
	}
}
