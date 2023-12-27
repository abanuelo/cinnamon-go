package main

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/abanuelo/cinnamon-go/cinnamon"
	"github.com/abanuelo/cinnamon-go/priorityq"
	"github.com/abanuelo/cinnamon-go/requestq"
	"google.golang.org/grpc"
)

// Total of 768 different priorities, initialized to middle
// 0-5 tiers where 0 > 5 and 0-127 cohorts where 0 > 127
var TIER_COHORT_THRESHOLD = 384

type CinnamonServiceServer struct {
	cinnamon.UnimplementedCinnamonServer
	PriorityQueue *priorityq.PriorityQueue
	mu            *sync.Mutex
}

func (s CinnamonServiceServer) Intercept(ctx context.Context, req *cinnamon.InterceptRequest) (*cinnamon.InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		s.mu.Lock()
		heap.Push(s.PriorityQueue, &priorityq.Item{Value: req.Route, Priority: int(req.Priority), Arrival: time.Now()})
		s.mu.Unlock()
		return &cinnamon.InterceptResponse{
			Accepted: true,
			Message:  s.PriorityQueue.PrintContents(),
		}, nil
	} else {
		result := fmt.Sprintf("Exceeds current threshold: %d", TIER_COHORT_THRESHOLD)
		return &cinnamon.InterceptResponse{
			Accepted: false,
			Message:  result,
		}, nil
	}
}

func main() {
	pq := make(priorityq.PriorityQueue, 0)
	var mutex sync.Mutex //Locking
	service := &CinnamonServiceServer{
		PriorityQueue: &pq,
		mu:            &mutex,
	}

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	grpcServer := grpc.NewServer()
	cinnamon.RegisterCinnamonServer(grpcServer, service)

	go func() {
		log.Println("gRPC server listening on :8089")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %s", err)
		}
	}()

	//TODO: add some locking mechanism around the priority queue
	numWorkers := 2
	// Wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(i)
		go requestq.Worker(i, &pq, &wg, &mutex)
	}

	// Start the goroutine for timeout of items in pq, setting it to a second for now
	maxAge := 1 * time.Second
	go priorityq.RemoveOldItems(&pq, maxAge, &mutex)

	// Keep the main goroutine running
	select {}
}
