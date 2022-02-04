package rpcs

import (
	"GolangImplementation/rpc-basic/db"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client() {
	fmt.Println("i am client")
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	c := db.NewGreetServiceClient(conn)
	fmt.Printf("create client %f\n", c)
}
