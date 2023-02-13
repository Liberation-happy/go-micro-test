package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"system/handler"
	pb "system/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "system"
	version = "latest"
)

func main() {
	consulReq := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"49.233.36.209:8500"}
	})
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReq),
	)
	srv.Init()

	// Register handler
	err := pb.RegisterSystemHandler(srv.Server(), new(handler.System))
	if err != nil {
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
