package cinnamon

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrianbrad/queue"
)

func writeToFile(item Item) {
	f, err := os.OpenFile("t.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(item)
	if err != nil {
		fmt.Println("Error with JSON marshal")
	}

	newLine := string(out)
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Worker(worker int, pq *queue.Priority[Item]) {
	for {
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			if pq.Size() > 0 {
				elem, err := pq.Get()
				if err != nil {
					fmt.Println("[ERROR] - Get() inside Worker for priority queue: ", err)
					return
				}
				if tq.Contains(elem) {
					fmt.Println("No action with dequeued item, has timed out")
					return
				}
				writeToFile(elem)
				CURR_INFLIGHT += 1
				// add to the inflight queue
				if err := iq.Offer(elem); err != nil {
					fmt.Println("[ERROR] - Offer() inside Worker inside inflight queue: ", err)
				}
				OUT += 1
			}
		}
	}
}
