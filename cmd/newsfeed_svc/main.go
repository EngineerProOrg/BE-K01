package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/internal/app/newsfeed_svc"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"google.golang.org/grpc"
)

var (
	path = flag.String("conf", "config.yml", "config path for this service")
)

func main() {
	// Start authenticate and post service

	conf, err := configs.GetNewsfeedConfig(*path)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	service, err := newsfeed_svc.NewNewsfeedService(conf)
	if err != nil {
		log.Fatalf("failed to init service %s", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	newsfeed.RegisterNewsfeedServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped %v", err)
	}
}
