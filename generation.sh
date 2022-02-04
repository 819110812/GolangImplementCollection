protoc -I=./protobufs --go_out=./protobufs ./protobufs/date.proto
protoc -I=./rpc-basic --go_out=./rpc-basic --go-grpc_out=./rpc-basic ./rpc-basic/greet.proto
protoc -I=./rpc-unary --go_out=./rpc-unary --go-grpc_out=./rpc-unary ./rpc-unary/greet.proto
