package cinnamon

import (
	"fmt"
	sync "sync"

	"github.com/abanuelo/cinnamon-go/queues"
)

func Worker(worker int, pq *queues.PriorityQueue, cq *queues.CircularQueue, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		mu.Lock()
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			if pq.Len() > 0 {
				req := pq.Dequeue()
				req.Status = int(Status_ERROR)
				CURR_INFLIGHT += 1
				fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
				cq.Enqueue(int(req.Priority))
				OUT += 1
				req.Status = int(Status_OK)
				mu.Unlock()
			} else {
				mu.Unlock()
			}
		} else {
			mu.Unlock()
		}
	}
}
