package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/abanuelo/cinnamon-go/cinnamon"
	"github.com/go-chi/chi"
)

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	// Attach the gRPC middleware to your Chi router
	r.Use(cinnamon.CinnamonMiddleware)

	// Add your other routes and middleware as needed
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(8 * time.Second)
		// fmt.Printf("HERE: Processed request: %s\n", "/hello")
		cinnamon.CURR_INFLIGHT -= 1
		w.Write([]byte("Hello"))
	})

	r.Get("/world", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		// fmt.Printf("HERE: Processed request: %s\n", "/world")
		cinnamon.CURR_INFLIGHT -= 1
		w.Write([]byte("World"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		// fmt.Printf("HERE: Processed request: %s\n", "/test")
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
