package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/abanuelo/cinnamon-go/cinnamon"
	"github.com/abanuelo/cinnamon-go/queues"
	"github.com/go-chi/chi"
)

func main() {
	pq := queues.NewPriorityQueue()
	cq := queues.NewCircularQueue(cinnamon.MAX_HISTORY)
	var mutex sync.Mutex //Lock shared across workers

	// Start the goroutine for timeout of items in pq, setting it to a second for now
	// TODO DO Not share Mutex here
	go queues.TimeoutItems(pq, cinnamon.MAX_AGE)

	// Running every 10 seconds to check if pq is still full to update threshold
	go cinnamon.LoadShed(pq, cq)

	// Wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines to pull items from priority queue
	for i := 0; i < cinnamon.NUM_WORKERS; i++ {
		wg.Add(i)
		go cinnamon.Worker(i, pq, cq, &wg, &mutex)
	}

	// Create a new Chi router
	r := chi.NewRouter()

	// Attach the gRPC middleware to your Chi router
	r.Use(cinnamon.CinnamonMiddleware(pq))

	// Add your other routes and middleware as needed
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Second)
		fmt.Printf("HERE: Processed request: %s\n", "/hello")
		cinnamon.CURR_INFLIGHT -= 1
		w.Write([]byte("Hello"))
	})

	r.Get("/world", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Printf("HERE: Processed request: %s\n", "/world")
		cinnamon.CURR_INFLIGHT -= 1
		w.Write([]byte("World"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Printf("HERE: Processed request: %s\n", "/test")
		cinnamon.CURR_INFLIGHT -= 1
		w.Write([]byte("Test"))
	})

	// Start the HTTP server
	addr := ":8089"
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Listen and serve
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s", addr)
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
