package main

import (
	"fmt"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-order/pkg/client"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-order/pkg/config"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-order/pkg/db"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-order/pkg/pb"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-order/pkg/services"
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

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
