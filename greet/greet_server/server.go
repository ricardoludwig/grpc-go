package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ricardoludwig/grpc-go-demo/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Função Greet foi invocada com %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Ola " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}
func main() {

	fmt.Println("Ola")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Falha na conexão: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha de conexão com o servidor: %v", err)
	}
}
