package main

import (
	rpcunary "GolangImplementation/rpc-unary"
	"GolangImplementation/rpc-unary/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ser := grpc.NewServer()
	db.RegisterGreetServiceServer(ser, &rpcunary.Server{})
	if err := ser.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
