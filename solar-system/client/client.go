package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ricardoludwig/grpc-go/solar-system/protos"
	"google.golang.org/grpc"
)

func main() {

	con := connect()
	defer con.Close()

	c := protos.NewSolarSystemServiceClient(con)
	req := &protos.SolarSystemRequest{
		Body: "Moon",
	}

	res, err := c.SolarSystem(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Solar System response: %v", res.GetBody())

}

func connect() (con *grpc.ClientConn) {

	fmt.Println("Starting Solar System Client")

	con, err := grpc.Dial("localhost:8180", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect:  %v", err)
	}
	return con
}
