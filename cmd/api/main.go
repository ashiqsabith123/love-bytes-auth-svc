package main

import (
	"net"

	"github.com/ashiqsabith123/auth-svc/pkg/config"
	"github.com/ashiqsabith123/auth-svc/pkg/di"
	"github.com/ashiqsabith123/auth-svc/pkg/utils"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"google.golang.org/grpc"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		logs.ErrLog.Fatal("Error while loading config", err)
	}

	err = logs.InitLogger("./pkg/logs/log.log")
	if err != nil {
		logs.ErrLog.Fatalln("Error while initilizing logger", err)
	}

	utils.InitTwilio(config)
	service := di.IntializeService(config)

	lis, err := net.Listen("tcp", config.Port.SvcPort)
	if err != nil {
		logs.ErrLog.Fatalln("Failed to listening:", err)
	}

	// credentials, err := helper.GetCertificates("cmd/cert/ca-cert.pem", "cmd/cert/server-cert.pem", "cmd/cert/server-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	logs.GenLog.Println("Auth Svc connected on", config.Port.SvcPort)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		logs.ErrLog.Fatalf("grpc serve err: %v", err)
	}

}
