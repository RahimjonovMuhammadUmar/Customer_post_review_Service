package api

import (
	// _ "exam/api_gateway/api/docs"
	v1 "exam/api_gateway/api/handler/v1"
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandleV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/customer", handlerV1.CreateCustomer)























	// api.DELETE("/product/:id", handlerV1.DeleteProduct)
	// api.PUT("/product", handlerV1.UpdateProduct)
	// api.GET("/product/:id", handlerV1.GetProduct)
	// api.POST("/type", handlerV1.CreateType)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
