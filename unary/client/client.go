package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ricardoludwig/grpc-go/unary/protos"
	"google.golang.org/grpc"
)

func main() {

	//con := connect()
	fmt.Println("Starting Unary Client")

	con, err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer con.Close()

	c := protos.NewHealthCheckServiceClient(con)
	req := &protos.HealthCheckRequest{
		Health: &protos.HealthCheck{
			Status: "Are you ok?",
		},
	}

	res, err := c.HealthCheck(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling health check: %v", err)
	}

	log.Printf("Health check server is: %v", res.GetResponse())
}

func connect() (con *grpc.ClientConn) {

	fmt.Println("Starting Unary Client")

	con, err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer con.Close()
	return con
}
