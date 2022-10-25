package main

import (
	"net"
	"exam/review_service/config"
	pbr "exam/review_service/genproto/review"
	"exam/review_service/pkg/db"
	"exam/review_service/pkg/logger"
	"exam/review_service/service"

	 "exam/review_service/service/grpc_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "rating_service")
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

	reviewService := service.NewReviewService(connDB, log, grpcClient)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while listening", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pbr.RegisterReviewServiceServer(s, reviewService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening %v", logger.Error(err))
	}

}
