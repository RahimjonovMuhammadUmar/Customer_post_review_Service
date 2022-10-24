package v1


import (
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}

// HandleV1Config ...
type HandleV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

// New ...
func New(h *HandleV1Config) *handlerV1 {
	return &handlerV1{
		log: h.Logger,
		serviceManager: h.ServiceManager,
		cfg: h.Cfg,
	}
}
