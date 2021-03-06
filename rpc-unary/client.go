package rpcunary

import (
	"GolangImplementation/rpc-unary/db"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Client() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	c := db.NewGreetServiceClient(cc)
	res := doUnaryClient(err, c)
	log.Println(res.Result)

}

func doUnaryClient(err error, c db.GreetServiceClient) *db.GreetResponse {
	req := &db.GreetRequest{
		Greeting: &db.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return res
}
