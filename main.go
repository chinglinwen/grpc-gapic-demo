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

func main() {
	grpcAddr := flag.String("grpc", "0.0.0.0:10001", "grpc addr")
	httpAddr := flag.String("http", "0.0.0.0:10000", "http addr")
	flag.Parse()

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	s := &services.ServerServices{}
	pb.RegisterServerServiceServer(grpcServer, s)

	mux := runtime.NewServeMux()
	err = pb.RegisterServerServiceHandlerServer(context.Background(), mux, s)
	if err != nil {
		log.Fatal(err)
	}

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", *httpAddr)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatal(http.ListenAndServe(*httpAddr, mux))
}
