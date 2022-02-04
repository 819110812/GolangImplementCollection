protoc -I=./protobufs --go_out=./protobufs ./protobufs/date.proto
protoc -I=./rpc --go_out=./rpc --go-grpc_out=./rpc ./rpc/greet.proto
