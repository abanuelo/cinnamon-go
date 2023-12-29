package cinnamon

import (
	"fmt"
	sync "sync"

	"github.com/abanuelo/cinnamon-go/queues"
	"github.com/adrianbrad/queue"
)

// func Worker(worker int, pq *queues.PriorityQueue, cq *queues.CircularQueue, wg *sync.WaitGroup, mu *sync.Mutex) {
// 	defer wg.Done()
// 	for {
// 		mu.Lock()
// 		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
// 			if pq.Len() > 0 {
// 				req := pq.Dequeue()
// 				req.Status = int(Status_OK)
// 				CURR_INFLIGHT += 1
// 				fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
// 				cq.Enqueue(int(req.Priority))
// 				OUT += 1
// 				mu.Unlock()
// 			} else {
// 				mu.Unlock()
// 			}
// 		} else {
// 			mu.Unlock()
// 		}
// 	}
// }

func Worker(worker int, pq *queue.Priority[queues.Item], wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for {
		// mu.Lock()
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			if pq.Size() > 0 {
				elem, err := pq.Get()
				if err != nil {
					// handle err
					fmt.Println("Error handling Get() inside Worker", err)
				}
				elem.Status = int(Status_OK)
				fmt.Println("Element dequeued and ready for work: ", elem)
				CURR_INFLIGHT += 1
				fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
				// cq.Enqueue(int(req.Priority))
				OUT += 1
				// mu.Unlock()
			} else {
				// mu.Unlock()
			}
		} else {
			// mu.Unlock()
		}
	}
}
