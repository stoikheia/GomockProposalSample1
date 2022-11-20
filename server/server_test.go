package main_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	pb "github.com/stoikheia/GomockProposalSample1/protobuf/helloworld"
	mock "github.com/stoikheia/GomockProposalSample1/protobuf/helloworld/mock"
	. "github.com/stoikheia/GomockProposalSample1/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestServe(t *testing.T) {
	tctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	expectReq := &pb.HelloRequest{Name: "Test"}
	expectedRes := &pb.HelloReply{Message: "Hello Test"}

	/*
		srv := &Server{} // no error when no mocking
	*/

	///*
	srv := mock.NewMockGreeterServer(gomock.NewController(t))
	srv.EXPECT().SayHello(gomock.Any(), expectReq).Return(expectedRes, nil)
	//*/

	blis := bufconn.Listen(1024 * 1024)
	s := Serve(srv, blis) // Error: missing method mustEmbedUnimplementedGreeterServer
	defer s.Stop()

	dialer := func(ctx context.Context, address string) (net.Conn, error) {
		return blis.Dial()
	}
	conn, err := grpc.DialContext(tctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	res, err := c.SayHello(tctx, expectReq)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	if diff := cmp.Diff(res, expectedRes, cmpopts.IgnoreUnexported(pb.HelloReply{})); diff != "" {
		t.Errorf("%s", diff)
	}

}
