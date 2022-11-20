package main

import (
	"context"
	"log"
	"net"

	pb "github.com/stoikheia/GomockProposalSample1/protobuf/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.GreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	name := in.GetName()
	log.Printf("Received: %v", name)
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Serve(srv pb.GreeterServer, lis net.Listener, opts ...grpc.ServerOption) *grpc.Server {
	s := grpc.NewServer(opts...)
	reflection.Register(s)
	//pb.RegisterGreeterServer(s, srv)
	s.RegisterService(&pb.Greeter_ServiceDesc, srv)
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return s
}
