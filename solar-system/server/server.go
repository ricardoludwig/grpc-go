package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/ricardoludwig/grpc-go/solar-system/protos"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	connect()
}

func connect() {

	fmt.Println("Starting Solar System Server...")

	list, err := net.Listen("tcp", "localhost:8180")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	conn := grpc.NewServer()
	protos.RegisterSolarSystemServiceServer(conn, &server{})
	if err := conn.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

var ErrorNotFound = errors.New("not found")

func (*server) SolarSystem(ctx context.Context, req *protos.SolarSystemRequest) (*protos.SolarSystemResponse, error) {

	body := req.GetBody()

	if body != "Earth" {
		return nil, fmt.Errorf("%v %w", body, ErrorNotFound)
	}

	response := &protos.SolarSystemResponse{
		Body: &protos.Body{
			Name: "Earth",
		},
	}
	return response, nil

}
