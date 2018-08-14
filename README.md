Demo GRPC client & server in Golang.

    go run server.go &

    go run client.go

How to run the protobuf compiler:

    protoc --go_out=plugins=grpc:double  double.proto
