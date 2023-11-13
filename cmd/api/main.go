package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/ashiqsabith123/auth-svc/pkg/di"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error while loading configz", err)
	}
	service := di.IntializeService(config)

	lis, err := net.Listen("tcp", config.Port.SvcPort)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}

	// credentials, err := helper.GetCertificates("cmd/cert/ca-cert.pem", "cmd/cert/server-cert.pem", "cmd/cert/server-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Auth Svc on", config.Port.SvcPort)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}

}
