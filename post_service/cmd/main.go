package main

import (
	"exam/post_service/config"
	pbp "exam/post_service/genproto/post"
	kafkaproducer "exam/post_service/kafka_producer"
	"exam/post_service/pkg/db"
	"exam/post_service/pkg/logger"
	"exam/post_service/service"
	"net"

	grpcClient "exam/post_service/service/grpc_client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Loglevel, "post_service")
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

	kafkaC, close, err := kafkaproducer.NewKafka(cfg)
	if err != nil {
		log.Fatal("Error while connecting to kafka", logger.Error(err))
	}
	defer close()

	postService := service.NewPostService(connDB, log, grpcClient, kafkaC)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while listening", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pbp.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening %v", logger.Error(err))
	}

}
