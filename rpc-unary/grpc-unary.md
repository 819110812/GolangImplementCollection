# grpc unary 
## 什么是grpc unary api?
grpc unary api是grpc的一种特殊的api，它只能调用一个rpc，并且只能调用一次。grpc unary api是通过
protobuf定义的，它的定义如下：
```go

```

## 定义一个grpc unary api
grpc unary api的定义如下：
```protobuf
syntax = "proto3";
package rpc_unary;

option go_package = "./db";

message Greeting {
  string first_name = 1;
  string last_name = 2;
}

// request
message GreetRequest {
  Greeting greeting = 1;
}

// response
message GreetResponse {
  string result = 1;
}

service GreetService {
  // unary api
  // Greet 是input
  // return GreetResponse 是output
  rpc Greet(GreetRequest) returns (GreetResponse) {}
}
```
使用命令行生成对应文件：
`protoc -I=./rpc-unary --go_out=./rpc-unary --go-grpc_out=./rpc-unary ./rpc-unary/greet.proto
`
可以发现在db目录下的greet_grpc.pd.go文件中多了很多的代码定义。

## 定义grpc unary server
在自动生成的代码中，greet server的接口定义如下：
```go
type GreetServiceServer interface {
	// unary api
	// Greet 是input
	// return GreetResponse 是output
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	mustEmbedUnimplementedGreetServiceServer()
}
```
因此我们定义的server需要继承这个接口，并且实现这个接口的方法。