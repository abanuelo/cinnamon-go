package cinnamon

import (
	context "context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	sync "sync"
	"time"

	"github.com/adrianbrad/queue"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Item struct {
	Priority int64
	Arrival  time.Time
	Status   int
}

func waitForAction(item Item, wg *sync.WaitGroup, qTime time.Time) {
	defer wg.Done()

	for !iq.Contains(item) {
		currentTime := time.Now()
		if currentTime.Sub(qTime) > MAX_AGE {
			break
		}
	}
}

func Intercept(ctx context.Context, req *InterceptRequest) (*InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		t := time.Now()
		item := Item{
			Priority: req.Priority,
			Arrival:  t,
			Status:   int(Status_PENDING),
		}
		if err := pq.Offer(item); err != nil {
			fmt.Println("Error inserting into priority queue with Offer(): ", err)
		}
		IN += 1

		var wg sync.WaitGroup
		wg.Add(1)
		go waitForAction(item, &wg, t)
		wg.Wait()

		if iq.Contains(item) {
			return &InterceptResponse{
				Accepted: true,
				Message:  "Processing request successfully",
			}, nil
		} else {
			// offer the item to timeout queue
			if err := tq.Offer(item); err != nil {
				// handle err
				fmt.Println("Error inserting into timeout queue with Offer(): ", err)
			}
			return nil, errors.New("Timeout in queue!")
		}

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
		_, err := Intercept(r.Context(), grpcRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to call main API: %v", err), http.StatusInternalServerError)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
	// }
}

var elems []Item

// priority queue
var pq = queue.NewPriority(
	elems,
	func(elem, otherElem Item) bool { return elem.Priority < otherElem.Priority },
)

// timeout queue
var tq = queue.NewCircular(elems, CIRCULAR_QUEUE_LENGTH)

// inflight queue
var iq = queue.NewCircular(elems, CIRCULAR_QUEUE_LENGTH)

func main() {
	fmt.Println("Inside main for middleware.go")

	// Running every 10 seconds - PID Controller
	go LoadShed(pq)

	// Start worker goroutines to pull items from priority queue
	for i := 0; i < NUM_WORKERS; i++ {
		go Worker(i, pq)
	}
}
