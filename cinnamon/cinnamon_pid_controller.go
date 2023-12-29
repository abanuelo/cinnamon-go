package cinnamon

import (
	"fmt"
	"time"

	"github.com/abanuelo/cinnamon-go/queues"
	"github.com/adrianbrad/queue"
)

// func LoadShed(pq *queues.PriorityQueue, cq *queues.CircularQueue) {
// 	ticker := time.NewTicker(10 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("Job: Checking Priority Queue...", time.Now())
// 			// Your job logic goes here
// 			// mu.Lock()
// 			if pq.Len() > 0 {
// 				fmt.Println("=========================================")
// 				fmt.Println("IN NEED OF PID CONTROLLER")
// 				fmt.Println(IN, OUT)
// 				fmt.Println(INFLIGHT_LIMIT, CURR_INFLIGHT)
// 				P := 0.0
// 				if OUT != 0 {
// 					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / OUT
// 				} else {
// 					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / INFLIGHT_LIMIT
// 				}

// 				fmt.Printf("P: %f\n", P)
// 				// if cq.CurrentCapacity() == MAX_HISTORY {
// 				cq.PrintQueue()
// 				newTreshold := cq.PercentileDistribution(P)
// 				fmt.Printf("NEW THRESHOLD: %d\n", newTreshold)
// 				// TIER_COHORT_THRESHOLD = newTreshold
// 				// }
// 				// pq.PrintContents()
// 				fmt.Println("=========================================")
// 				// TODO CHANGE THIS ONCE WE KNOW IT WORKS
// 				// IN = 0.0
// 				// OUT = 0.0
// 			}
// 			// mu.Unlock()
// 		}
// 	}
// }

func LoadShed(pq *queue.Priority[queues.Item]) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Job: Checking Priority Queue...", time.Now())
			// Your job logic goes here
			// mu.Lock()
			if pq.Size() > 0 {
				fmt.Println("=========================================")
				fmt.Println("IN NEED OF PID CONTROLLER")
				fmt.Println(IN, OUT)
				fmt.Println(INFLIGHT_LIMIT, CURR_INFLIGHT)
				P := 0.0
				if OUT != 0 {
					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / OUT
				} else {
					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / INFLIGHT_LIMIT
				}

				fmt.Printf("P: %f\n", P)
				// if cq.CurrentCapacity() == MAX_HISTORY {
				// cq.PrintQueue()
				// newTreshold := cq.PercentileDistribution(P)
				// fmt.Printf("NEW THRESHOLD: %d\n", newTreshold)
				// TIER_COHORT_THRESHOLD = newTreshold
				// }
				// pq.PrintContents()
				fmt.Println("=========================================")
				// TODO CHANGE THIS ONCE WE KNOW IT WORKS
				// IN = 0.0
				// OUT = 0.0
			}
			// mu.Unlock()
		}
	}
}
