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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func waitForChange(item *queues.Item, wg *sync.WaitGroup) {
	defer wg.Done()

	for item.Status == int(Status_PENDING) {
		time.Sleep(100 * time.Millisecond)
	}
}

func Intercept(ctx context.Context, req *InterceptRequest, pq *queues.PriorityQueue) (*InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		item := queues.Item{
			Method:   req.Method,
			Url:      req.Url,
			Priority: req.Priority,
			Arrival:  time.Now(),
			Status:   int(Status_PENDING),
		}
		//TODO check if we even need the Mutex lock
		// mu.Lock()
		pq.Enqueue(&item)
		// heap.Push(pq, &item)
		IN += 1
		// mu.Unlock()
		var wg sync.WaitGroup
		wg.Add(1)
		go waitForChange(&item, &wg)

		wg.Wait()

		fmt.Println(item)

		if item.Status == int(Status_OK) {
			return &InterceptResponse{
				Accepted: true,
				Message:  "Processing request successfully",
			}, nil
		}
		// } else if item.Status == int(Status_ERROR) {
		return nil, errors.New("Timeout in queue!")
		// }

	} else {
		result := fmt.Sprintf("Exceeds current threshold: %d", TIER_COHORT_THRESHOLD)
		return nil, errors.New(result)
	}
}

// func CinnamonMiddleware(next http.Handler, pq *queues.PriorityQueue) http.Handler {
func CinnamonMiddleware(pq *queues.PriorityQueue) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path[len(r.URL.Path)-1:] == "*" {
				// If it does, skip the middleware and call the next handler directly
				next.ServeHTTP(w, r)
				return
			}
			// Seed the random number generator with the current time
			rand.Seed(time.Now().UnixNano())

			// Generate a random integer between 0 and 767
			randomInteger := rand.Intn(768)

			// Create a gRPC request
			grpcRequest := &InterceptRequest{
				Method:   r.Method,
				Url:      r.URL.String(),
				Priority: int64(randomInteger),
				Arrival:  timestamppb.New(time.Now()),
				Status:   *Status_PENDING.Enum(),
			}

			// Call the gRPC method on the main API
			grpcResponse, err := Intercept(r.Context(), grpcRequest, pq)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to call main API: %v", err), http.StatusInternalServerError)
				return
			}
			fmt.Println(grpcResponse)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
