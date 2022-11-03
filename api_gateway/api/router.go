package api

import (
	_ "exam/api_gateway/api/docs"
	v1 "exam/api_gateway/api/handler/v1"
	"exam/api_gateway/api/middleware"
	"exam/api_gateway/api/token"
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
	"exam/api_gateway/storage/repo"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Redis          repo.InMemoryStorageI
	CasbinEnforcer        *casbin.Enforcer
}


// New ...
// @title           exam api
// @version         1.0
// @description     This is exam server api server
// @termsOfService  2 term exam

// @contact.name   Muhammad Umar
// @contact.url    https://t.me/muhammad_ummar
// @contact.email  rahimzanovmuhammadumar@gmail.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandleV1Config{
		Logger:          option.Logger,
		ServiceManager:  option.ServiceManager,
		Cfg:             option.Conf,
		InMemoryStorage: option.Redis,
	})

	jwtHandler := token.JWTHandler{
		SignInKey: option.Conf.SignInKey,
		Log:       option.Logger,
	}
	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, config.Load()))
	api := router.Group("/v1")

	//customer
	api.POST("/customer", handlerV1.CreateCustomer)
	api.GET("/customer/:id", handlerV1.GetCustomer)
	api.PUT("/customer", handlerV1.UpdateCustomer)
	api.DELETE("/customer/:id", handlerV1.DeleteCustomer)
	api.GET("/customer/search", handlerV1.SearchCustomer)

	//post
	api.POST("/post", handlerV1.CreatePost)
	api.GET("/post/:id", handlerV1.GetPostWithCustomerInfo)
	api.GET("/post/customers_posts/:id", handlerV1.GetPostsOfCustomer)
	api.PUT("/post", handlerV1.UpdatePost)
	api.DELETE("/post/:id", handlerV1.DeletePost)
	api.DELETE("/post/delete_customers_posts/:id", handlerV1.DeletePostByCustomerId)

	// review
	api.POST("/review", handlerV1.CreateReview)
	api.DELETE("/review/:id", handlerV1.DeleteReview)
	api.GET("/review/:id", handlerV1.GetReview)
	api.DELETE("/review_by_custID/:id", handlerV1.DeleteCustomerRates)

	// register
	api.POST("/register", handlerV1.RegisterCustomer)
	api.GET("/register/:code/:email", handlerV1.VerifyRegistration)
	api.GET("/login/:email/:password", handlerV1.Login)

	// admin
	api.GET("/admin/login/:username/:password", handlerV1.AdminLogin)

	// moderator
	api.GET("/moderator/login/:username/:password", handlerV1.ModeratorLogin)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
