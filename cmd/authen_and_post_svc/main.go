package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EngineerProOrg/BE-K01/internal/app/authen_and_post_svc"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"google.golang.org/grpc"
)

func main() {
	// Start authenticate and post service
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", "1080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := authen_and_post_svc.NewAuthenticateAndPostService()
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	authen_and_post.RegisterAuthenticateAndPostServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}
}
