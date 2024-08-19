package main

import (
	"fmt"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-product/pkg/config"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-product/pkg/db"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-product/pkg/pb"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-product/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := service.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
