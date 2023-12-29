package cinnamon

import (
	context "context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	sync "sync"
	"time"

	"github.com/abanuelo/cinnamon-go/queues"
	"github.com/adrianbrad/queue"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func waitForChange(item queues.Item, wg *sync.WaitGroup) {
	defer wg.Done()

	for pq.Contains(item) {
		fmt.Println("Waiting: .... but here is the queue size: ", pq.Size())
		time.Sleep(500 * time.Millisecond)
	}

	// for item.Status == int(Status_PENDING) {
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("item: ", item)
	// }
	fmt.Println("item dequeued: ", item)
}

func Intercept(ctx context.Context, req *InterceptRequest) (*InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		item := queues.Item{
			Priority: req.Priority,
			Arrival:  time.Now(),
			Status:   int(Status_PENDING),
		}
		if err := pq.Offer(item); err != nil {
			// handle err
			fmt.Println("Error inserting into queue with Offer(): ", err)
		}
		IN += 1

		var wg sync.WaitGroup
		wg.Add(1)
		go waitForChange(item, &wg)
		wg.Wait()

		fmt.Println(item)

		// TODO create a timeout queue and/or figure out how to status in item
		// if item.Status == int(Status_OK) {
		return &InterceptResponse{
			Accepted: true,
			Message:  "Processing request successfully",
		}, nil
		// }
		// } else if item.Status == int(Status_ERROR) {
		// return nil, errors.New("Timeout in queue!")
		// }

	} else {
		result := fmt.Sprintf("Exceeds current threshold: %d", TIER_COHORT_THRESHOLD)
		return nil, errors.New(result)
	}
}

// func CinnamonMiddleware(pq *queues.PriorityQueue) func(next http.Handler) http.Handler {
// return func(next http.Handler) http.Handler {
func CinnamonMiddleware(next http.Handler) http.Handler {
	var pqOnce sync.Once
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[len(r.URL.Path)-1:] == "*" {
			// If it does, skip the middleware and call the next handler directly
			next.ServeHTTP(w, r)
			return
		}
		pqOnce.Do(func() {
			main()
		})

		// Generate a random integer between 0 and 767
		tierOne := rand.Intn(127)
		tierFour := rand.Intn(639-511+1) + 511

		var selectedValue int
		if rand.Intn(2) == 0 {
			selectedValue = tierOne
		} else {
			selectedValue = tierFour
		}

		// Create a gRPC request
		grpcRequest := &InterceptRequest{
			Method:   r.Method,
			Url:      r.URL.String(),
			Priority: int64(selectedValue),
			Arrival:  timestamppb.New(time.Now()),
			Status:   *Status_PENDING.Enum(),
		}

		// Call the gRPC method on the main API
		grpcResponse, err := Intercept(r.Context(), grpcRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call main API: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Println(grpcResponse)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
	// }
}

var elems []queues.Item

var pq = queue.NewPriority(
	elems,
	func(elem, otherElem queues.Item) bool { return elem.Priority < otherElem.Priority },
)

// timeout queue
// var tq = queue.(elems, 0)

// var cq = queues.NewCircularQueue(MAX_HISTORY)

func main() {
	// pq := queues.NewPriorityQueue()
	// cq := queues.NewCircularQueue(MAX_HISTORY)
	fmt.Println("Inside main for middleware.go")
	// Start the goroutine for timeout of items in pq, setting it to a second for now
	// TODO DO Not share Mutex here
	go queues.TimeoutItems(pq, MAX_AGE)

	// Running every 10 seconds to check if pq is still full to update threshold
	// go LoadShed(pq, cq)
	go LoadShed(pq)

	// Wait group to wait for all workers to finish
	var wg sync.WaitGroup
	var mutex sync.Mutex //Lock shared across workers
	// Start worker goroutines to pull items from priority queue
	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(i)
		// go Worker(i, pq, cq, &wg, &mutex)
		go Worker(i, pq, &wg, &mutex)
	}
}
