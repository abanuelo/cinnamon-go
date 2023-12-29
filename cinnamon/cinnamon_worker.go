package cinnamon

import (
	"encoding/json"
	"fmt"
	"os"
	sync "sync"
	"time"

	"github.com/adrianbrad/queue"
)

func writeToFile(item Item) {
	f, err := os.OpenFile("priority_queue.txt", os.O_APPEND|os.O_WRONLY, 0644)
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
	// fmt.Println("file appended successfully")
}

func Worker(worker int, pq *queue.Priority[Item], wg *sync.WaitGroup) {
	// defer wg.Done()
	for {
		if CURR_INFLIGHT+1 < INFLIGHT_LIMIT {
			if pq.Size() > 0 {
				elem, err := pq.Get()
				if err != nil {
					fmt.Println("[ERROR] - Get() inside Worker for priority queue: ", err)
					return
				}
				writeToFile(elem)
				CURR_INFLIGHT += 1
				// add to the inflight queue
				if err := iq.Offer(elem); err != nil {
					fmt.Println("[ERROR] - Offer() inside Worker inside inflight queue: ", err)
				}
				OUT += 1
				// fmt.Printf("Current CURR_INFLIGHT: %f with worker id: %d\n", CURR_INFLIGHT, worker)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
