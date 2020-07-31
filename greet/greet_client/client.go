package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ricardoludwig/grpc-go-demo/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
		os.Exit(1)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Unary RPC ...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Ludwig",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro enquanto chamava Greet RPC: %v", err)
		os.Exit(1)
	}
	log.Printf("Resposta de Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Iniciando o Server Streaming RPC ...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ricardo",
			LastName:  "Ludwig",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error enquanto chamava GreeManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//Terminou o Stream
			fmt.Println("Terminou o stream")
			break
		}
		if err != nil {
			log.Fatalf("Erro enquanto lia o stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
