package main

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/abanuelo/cinnamon-go/cinnamon"
	"github.com/abanuelo/cinnamon-go/priorityq"
	"google.golang.org/grpc"
)

type myCinnamonServer struct {
	cinnamon.UnimplementedInvoicerServer
}

func (s myCinnamonServer) Create(ctx context.Context, req *cinnamon.CreateRequest) (*cinnamon.CreateResponse, error) {
	return &cinnamon.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

// Total of 768 different priorities, initialized to middle
// 0-5 tiers where 0 > 5 and 0-127 cohorts where 0 > 127
var TIER_COHORT_THRESHOLD = 384

func main() {
	pq := make(priorityq.PriorityQueue, 0)

	heap.Push(&pq, &priorityq.Item{Value: "Item 1", Priority: 3})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*priorityq.Item)
		fmt.Printf("Popped: %s (Priority: %d)\n", item.Value, item.Priority)
	}

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myCinnamonServer{}
	cinnamon.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Impossible to service: %s", err)
	}
}
