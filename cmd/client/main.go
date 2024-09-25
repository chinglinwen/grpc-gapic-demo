package main

import (
	"context"
	"flag"
	"log"

	pb "core.wcloud.io/generated/grpcgen" // Update
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcAddr := flag.String("grpc", "localhost:10001", "grpc addr")
	flag.Parse()
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(*grpcAddr, opts)
	if err != nil {
		log.Fatalf("failed to new conn, err: %v", err)
	}
	defer conn.Close()

	cli := pb.NewServerServiceClient(conn)
	resp, err := cli.ListServers(context.Background(), &pb.ListServersRequest{})
	if err != nil {
		log.Fatalf("failed list servers, err: %v", err)
	}
	log.Printf("servers: %+v", resp.GetServers())
}
