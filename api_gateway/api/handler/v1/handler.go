package v1

import (
	t "exam/api_gateway/api/token"
	"exam/api_gateway/config"
	"exam/api_gateway/pkg/logger"
	"exam/api_gateway/services"
	"exam/api_gateway/storage/repo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type handlerV1 struct {
	log             logger.Logger
	serviceManager  services.IServiceManager
	InMemoryStorage repo.InMemoryStorageI
	cfg             config.Config
	jwthandler      t.JWTHandler
}

// HandleV1Config ...
type HandleV1Config struct {
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	InMemoryStorage repo.InMemoryStorageI
	Cfg             config.Config
	JWTHandler t.JWTHandler
}

// New ...
func New(h *HandleV1Config) *handlerV1 {
	return &handlerV1{
		log:             h.Logger,
		serviceManager:  h.ServiceManager,
		InMemoryStorage: h.InMemoryStorage,
		cfg:             h.Cfg,
		jwthandler: h.JWTHandler,
	}
}



func GetClaims(h handlerV1, c *gin.Context) (*t.CustomClaims, error) {
	claims := t.CustomClaims{}
	strToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte("UmarSecret"), nil })
	if err != nil {
		h.log.Error("invalid access token")
		return nil, err
	}

	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Exp = rawClaims["exp"].(float64)

	aud := cast.ToStringSlice(rawClaims["aud"])
	claims.Aud = aud
	claims.Role = rawClaims["role"].(string)
	claims.Sub = rawClaims["sub"].(string)
	claims.Token = token
	return &claims, nil
}