package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EngineerProOrg/BE-K01/internal/app/newsfeed_svc"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"google.golang.org/grpc"
)

func main() {
	// Start authenticate and post service
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", "1080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := newsfeed_svc.NewNewsfeedService()
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	newsfeed.RegisterNewsfeedServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}
}
