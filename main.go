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
var NUM_WORKERS = 2
var MAX_AGE = 1 * time.Second

type CinnamonServiceServer struct {
	cinnamon.UnimplementedCinnamonServer
	PriorityQueue *priorityq.PriorityQueue
	mu            *sync.Mutex
}

func waitForChange(item *priorityq.Item, wg *sync.WaitGroup) {
	defer wg.Done()

	for item.Processed != "processed" && item.Processed != "timeout" {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Sleeping...still at no")
	}

	fmt.Println("item.Processed value changed")
	fmt.Println(item.Processed)
}

func (s CinnamonServiceServer) Intercept(ctx context.Context, req *cinnamon.InterceptRequest) (*cinnamon.InterceptResponse, error) {
	if int(req.Priority) <= TIER_COHORT_THRESHOLD {
		item := priorityq.Item{Value: req.Route, Priority: int(req.Priority), Arrival: time.Now(), Processed: "no"}
		s.mu.Lock()
		heap.Push(s.PriorityQueue, &item)
		s.mu.Unlock()

		var wg sync.WaitGroup
		wg.Add(1)
		go waitForChange(&item, &wg)
		wg.Wait()

		accepted := true
		message := ""
		if item.Processed == "processed" {
			accepted = true
			message = "Processed"
		} else if item.Processed == "timeout" {
			accepted = false
			message = "Timed Out!"
		}

		return &cinnamon.InterceptResponse{
			Accepted: accepted,
			Message:  message,
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

	// Wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(i)
		go requestq.Worker(i, &pq, &wg, &mutex)
	}

	// Start the goroutine for timeout of items in pq, setting it to a second for now
	go priorityq.RemoveOldItems(&pq, MAX_AGE, &mutex)

	// Keep the main goroutine running
	select {}
}
