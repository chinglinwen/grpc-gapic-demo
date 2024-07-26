package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	pb "core.wcloud.io/generated/grpcgen"
	"core.wcloud.io/server/services"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	s := &services.ServerServices{}
	pb.RegisterServerServiceServer(grpcServer, s)

	mux := runtime.NewServeMux()
	err = pb.RegisterServerServiceHandlerServer(context.Background(), mux, s)
	if err != nil {
		return err
	}

	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	return http.ListenAndServe(":11000", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
