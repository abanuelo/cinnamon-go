package cinnamon

import (
	"fmt"
	"time"

	"github.com/adrianbrad/queue"
)

type Item struct {
	Priority int64
	Arrival  time.Time
	Status   int
}

func TimeoutItems(pq *queue.Priority[Item], maxAge time.Duration) {
	// TODO REMOVE THE NEED TO DO THIS EVERY 4 SECONDS, OR FIND AN INTERVAL IN ORDER TO REMOVE ITEMS FROM QUEUE
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if pq.Size() > 0 {
				elem, err := pq.Peek()
				if err != nil {
					fmt.Println("[ERROR] - Peek() inside TimeoutItems for priority queue", err)
					return
				}
				currentTime := time.Now()
				if currentTime.Sub(elem.Arrival) > maxAge {
					elem, err := pq.Get()
					if err != nil {
						fmt.Println("[ERROR] - Get() inside TimeoutItems for priority queue", err)
						return
					}
					// add to the inflight queue
					if err := tq.Offer(elem); err != nil {
						fmt.Println("[ERROR] - Offer() inside Worker inside timeout queue: ", err)
					}
				}
			}
		}
	}
}
