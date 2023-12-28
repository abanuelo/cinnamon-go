package cinnamon

import (
	"container/heap"
	context "context"
	"errors"
	"fmt"
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
		// fmt.Println("Sleeping...still at no")
	}
	// fmt.Println("item.Status value changed")
	// fmt.Println(item.Status)
}

func Intercept(ctx context.Context, req *InterceptRequest, pq *queues.PriorityQueue, mu *sync.Mutex) (*InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		item := queues.Item{
			Method:   req.Method,
			Url:      req.Url,
			Priority: req.Priority,
			Arrival:  time.Now(),
			Status:   int(Status_PENDING),
		}
		//TODO check if we even need the Mutex lock
		mu.Lock()
		heap.Push(pq, &item)
		IN += 1
		mu.Unlock()

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
func CinnamonMiddleware(pq *queues.PriorityQueue, mu *sync.Mutex) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path[len(r.URL.Path)-1:] == "*" {
				// If it does, skip the middleware and call the next handler directly
				next.ServeHTTP(w, r)
				return
			}

			// Set up a gRPC client connection to the main API
			// grpcConn, err := grpc.Dial(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
			// if err != nil {
			// 	http.Error(w, "Failed to connect to main API", http.StatusInternalServerError)
			// 	return
			// }
			// defer grpcConn.Close()

			// Create a gRPC request
			// TODO need to change Priority to be intercepted from HTTP request
			grpcRequest := &InterceptRequest{
				Method:   r.Method,
				Url:      r.URL.String(),
				Priority: int64(TIER_COHORT_THRESHOLD),
				Arrival:  timestamppb.New(time.Now()),
				Status:   *Status_PENDING.Enum(),
			}
			// Call the gRPC method on the main API
			grpcResponse, err := Intercept(r.Context(), grpcRequest, pq, mu)
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
