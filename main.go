package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	// "github.com/googleapis/gapic-showcase/server/services"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	pb "core.wcloud.io/generated/grpcgen" // Update
	"core.wcloud.io/server/services"      // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	// s := services.Backend{}
	s := &services.ServerServices{}
	pb.RegisterServerServiceServer(grpcServer, s)

	mux := runtime.NewServeMux()
	err := pb.RegisterServerServiceHandlerServer(ctx, mux, s)
	if err != nil {
		return err
	}

	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, h2s),
	}
	return h1s.ListenAndServe()
}

func run2() error {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	// s := services.Backend{}
	s := &services.ServerServices{}
	pb.RegisterServerServiceServer(grpcServer, s)

	mux := runtime.NewServeMux()
	err = pb.RegisterServerServiceHandlerServer(context.Background(), mux, s)
	if err != nil {
		return err
	}

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	// // Register gRPC server endpoint
	// // Note: Make sure the gRPC server is running properly and accessible
	// mux := runtime.NewServeMux()
	// // opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err = pb.RegisterLibraryServiceHandlerServer(context.Background(), mux, ls)
	// if err != nil {
	// 	return err
	// }

	// // Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":11000", mux)
}

func main() {
	flag.Parse()

	if err := run2(); err != nil {
		grpclog.Fatal(err)
	}
}
