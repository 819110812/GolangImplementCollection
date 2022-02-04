package rpcs

import (
	"GolangImplementation/rpc-basic/db"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	db.UnimplementedGreetServiceServer
}

func RpcServer() {
	fmt.Println("Hello, 世界")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}

	ser := grpc.NewServer()
	db.RegisterGreetServiceServer(ser, &server{})
	if err := ser.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
