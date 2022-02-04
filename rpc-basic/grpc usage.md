# Golang grpc 使用
## 通过.proto生成文件
在新版的protobuf的命令行中，删除了--plugin的写法，改用--go-grpc_out来生成grpc文件。
`protoc -I=./rpc --go_out=./rpc --go-grpc_out=./rpc ./rpc/greet.proto
`

## 创建server端
```go
type server struct {
	// 定义一个结构体，用于存储服务端的数据
	// 必须加上以下这段UnimplementedGreetServiceServer来继承greetpb.GreetServiceServer
	db.UnimplementedGreetServiceServer
}

func RpcServer() {
	fmt.Println("Hello, 世界")
	// 监听一个端口
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}

	ser := grpc.NewServer()
	// 注册grpc服务
	db.RegisterGreetServiceServer(ser, &server{})
	if err := ser.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
```

## 创建client端

```go
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
```