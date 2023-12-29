package queues

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
	// TODO REMOVE THE NEED TO DO THIS EVERY 4 SECONDS
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// fmt.Println("PQ AT THE MOMENT: ")
			// pq.PrintContents()
			currentTime := time.Now()

			elem, err := pq.Peek()
			if err != nil {
				fmt.Println("Error handling Peek()", err)
			}

			if pq.Size() > 0 && currentTime.Sub(elem.Arrival) > maxAge {
				fmt.Println("****************************************")
				fmt.Println("Current time: ", currentTime)
				fmt.Println("Arrival time: ", elem.Arrival)
				fmt.Println("Current time - Arrival time: ", currentTime.Sub(elem.Arrival))
				fmt.Println("Max Age: ", maxAge)
				fmt.Println("****************************************")

				elem, err := pq.Get()
				if err != nil {
					// handle err
					fmt.Println("Error handling Get()", err)
				}
				elem.Status = 1
			}
		}
	}
}
