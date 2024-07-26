package services

import (
	"context"

	// pb "github.com/googleapis/gapic-showcase/server/genproto"
	pb "core.wcloud.io/generated/grpcgen"
	// locpb "google.golang.org/genproto/googleapis/cloud/location"
	// iampb "google.golang.org/genproto/googleapis/iam/v1"
)

type ServerServices struct {
	pb.ServerServiceServer
}

func (s *ServerServices) Listservers(context.Context, *pb.ListServersRequest) (*pb.ListServersResponse, error) {
	return &pb.ListServersResponse{
		Servers: []*pb.Server{
			{
				Name: "s1",
			},
			{
				Name: "s2",
			},
		},
	}, nil
}
