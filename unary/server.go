package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ricardoludwig/grpc-go/unary/protos"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) HealthCheck(ctx context.Context, req *protos.HealthCheckRequest) (*protos.HealthCheckResponse, error) {
	fmt.Printf("Teste")
	res := &protos.HealthCheckResponse{
		Response: "OK",
	}
	return res, nil
}

func main() {

	fmt.Println("Starting Unary Server")

	list, err := net.Listen("tcp", "0.0.0.0:8180")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterHealthCheckServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
