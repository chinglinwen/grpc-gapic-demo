package services

import (
	"log"

	"github.com/googleapis/gapic-showcase/server"
	// pb "github.com/googleapis/gapic-showcase/server/genproto"
	pb "core.wcloud.io/generated/grpcgen"

	// locpb "google.golang.org/genproto/googleapis/cloud/location"
	// iampb "google.golang.org/genproto/googleapis/iam/v1"
	lropb "google.golang.org/genproto/googleapis/longrunning"
)

// Backend contains the various service backends that will be
// accessible via one or more transport endpoints.
type Backend struct {
	// Showcase schema
	// EchoServer            pb.EchoServer
	// IdentityServer        pb.IdentityServer
	// MessagingServer       pb.MessagingServer
	// SequenceServiceServer pb.SequenceServiceServer
	// ComplianceServer      pb.ComplianceServer
	// TestingServer         pb.TestingServer
	ServerServiceServer pb.ServerServiceServer

	// Supporting protos
	OperationsServer lropb.OperationsServer
	// LocationsServer  locpb.LocationsServer
	// IAMPolicyServer  iampb.IAMPolicyServer

	// Other supporting data structures
	StdLog, ErrLog   *log.Logger
	ObserverRegistry server.GrpcObserverRegistry
}
