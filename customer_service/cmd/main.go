package main

import (
	"exam/customer_service/config"
	pbc "exam/customer_service/genproto/customer"
	"exam/customer_service/pkg/db"
	"exam/customer_service/pkg/logger"
	"exam/customer_service/service"
	"net"

	grpcClient "exam/customer_service/service/grpc_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Loglevel, "customer_service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("error while connDB, err := db.ConnectToDB(cfg)", logger.Error(err))
	}

	grpcClient, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("error while grpcClient, err := grpcClient.New(cfg)", logger.Error(err))
	}

	customerService := service.NewCustomerService(connDB, log, grpcClient)
	lis, err := net.Listen("172.17.0.3", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while listening", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pbc.RegisterCustomerServiceServer(s, customerService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening %v", logger.Error(err))
	}

}
