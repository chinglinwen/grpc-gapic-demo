package main

import (
	"context"
	"log"
	"testing"

	pb "core.wcloud.io/generated/grpcgen" // Update
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPC(t *testing.T) {
	cc, err := grpc.NewClient("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("new client err: %v", err)
	}
	cli := pb.NewServerServiceClient(cc)
	r, err := cli.Listservers(context.Background(), &pb.ListServersRequest{})
	if err != nil {
		log.Fatalf("list err: %v", err)
	}
	log.Printf("resp: %+v", r)
}
