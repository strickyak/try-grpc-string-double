package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/strickyak/try-grpc-string-double/double"
	"google.golang.org/grpc"
)

type doubleServer struct{}

func (s *doubleServer) DoubleString(ctx context.Context, in *pb.DoubleStringRequest) (*pb.DoubleStringReply, error) {
	doubled := in.Input + in.Input
	reply := &pb.DoubleStringReply{
		Output: doubled,
	}
	return reply, nil
}

var bind = flag.String("bind", "localhost:8080", "Port to bind to")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Panicf("Cannot bind to %s: %s", *bind, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDoubleServiceServer(grpcServer, &doubleServer{})
	grpcServer.Serve(lis)
}
