package main

import (
	"exam/api_gateway/api"
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
	r "exam/api_gateway/storage/redis"
	"log"

	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"

	// gormadapter "github.com/casbin/gorm-adapter/v2"

	"github.com/casbin/casbin/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	tracecfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831", // replace host
		},
	}

	closer, err := tracecfg.InitGlobalTracer(
		"api-gateway",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}

	var casbinEnforcer *casbin.Enforcer
	cfg := config.Load()
	log := logger.New(cfg.Loglevel, "api_gateway")

	// psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresDB,
	// )
	// enf, err := gormadapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("gorm adapter error", logger.Error(err))
	// 	return
	// }
	casbinEnforcer, err = casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CsvFilePath)
	if err != nil {
		log.Error("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("serviceManager, err := services.NewServiceManager(&cfg)", logger.Error(err))
	}
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis:6379")
		},
	}

	server := api.New(api.Option{
		Logger:         log,
		Conf:           cfg,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(pool),
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
