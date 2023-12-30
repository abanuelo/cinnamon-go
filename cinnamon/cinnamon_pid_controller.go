package cinnamon

import (
	"fmt"
	"time"

	"github.com/adrianbrad/queue"
)

func LoadShed(pq *queue.Priority[Item]) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Job: Checking Priority Queue...", time.Now())
			if pq.Size() > 0 {
				fmt.Println("=========================================")
				fmt.Println("IN NEED OF PID CONTROLLER")
				fmt.Println("IN: ", IN)
				fmt.Println("OUT: ", OUT)
				fmt.Println("INFLIGHT_LIMIT: ", INFLIGHT_LIMIT)
				fmt.Println("CURR_INFLIGHT: ", CURR_INFLIGHT)
				P := 0.0
				if OUT != 0 {
					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / OUT
				} else {
					P = (IN - (OUT + (INFLIGHT_LIMIT - CURR_INFLIGHT))) / INFLIGHT_LIMIT
				}

				fmt.Printf("P: %f\n", P)
				// TODO CALCULATE CDF of last 1000 seen items in IQ
				// TODO Update threshold on that calculation
				fmt.Println("=========================================")
				// Update IN and OUT to zero
				// TODO CHANGE THIS ONCE WE KNOW IT WORKS
				// IN = 0.0
				// OUT = 0.0
			}
		}
	}
}
