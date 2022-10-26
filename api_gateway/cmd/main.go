package main

import (
	"exam/api_gateway/api"
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
	r "exam/api_gateway/storage/redis"

	"github.com/gomodule/redigo/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.Loglevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("serviceManager, err := services.NewServiceManager(&cfg)", logger.Error(err))
	}

	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	server := api.New(api.Option{
		Logger:         log,
		Conf:           cfg,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(pool),
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
