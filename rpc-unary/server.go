package rpcunary

import (
	"GolangImplementation/rpc-unary/db"
	"context"
	"fmt"
)

type Server struct {
	db.UnimplementedGreetServiceServer
}

func (s *Server) Greet(context context.Context, request *db.GreetRequest) (*db.GreetResponse, error) {
	firstname := request.GetGreeting().GetFirstName()
	fmt.Println("Greeting received: ", firstname)
	return &db.GreetResponse{
		Result: "Hello" + firstname,
	}, nil
}
