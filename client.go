package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/strickyak/try-grpc-string-double/double"
	"google.golang.org/grpc"
)

var port = flag.String("conn", "localhost:8080", "Port to connect to")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*port, grpc.WithInsecure())
	if err != nil {
		log.Panicf("Cannot connect to %s: %s", *port, err)
	}
	defer conn.Close()
	client := pb.NewDoubleServiceClient(conn)

	req := &pb.DoubleStringRequest{
		Input: "trouble",
	}
	res, err := client.DoubleString(context.TODO(), req)
	if err != nil {
		log.Panicf("Cannot double:  %s", err)
	}
	println(res.Output)
}
