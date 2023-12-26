package main

import (
	"context"
	"log"
	"net"

	"github.com/abanuelo/cinnamon-go/cinnamon"
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

func main() {
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
