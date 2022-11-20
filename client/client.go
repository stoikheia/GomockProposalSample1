package main

import (
	"time"
	"log"
	"context"
	pb "github.com/stoikheia/GomockProposalSample1/protobuf/helloworld"
)

func SendUnary(c pb.GreeterClient) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Demo"})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}