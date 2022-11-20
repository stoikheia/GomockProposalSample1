package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := &Server{}
	Serve(srv, lis)
}
