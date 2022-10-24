package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string // develop, staging, production
	Loglevel          string
	HTTPPort           string
	
	CustomerServiceHost  string
	CustomerServicePort  int
	PostServiceHost  string
	PostServicePort  int
	ReviewServiceHost  string
	ReviewServicePort  int
	CtxTimeout int
}

// Load loads environmen vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.CustomerServiceHost = cast.ToString(getOrReturnDefault("CUSTOMER_SERVICE_HOST", "localhost"))
	c.CustomerServicePort = cast.ToInt(getOrReturnDefault("CUSTOMER_SERVICE_PORT", 9000))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 9090))

	c.ReviewServiceHost = cast.ToString(getOrReturnDefault("REVIEW_SERVICE_HOST", "localhost"))
	c.ReviewServicePort = cast.ToInt(getOrReturnDefault("REVIEW_SERVICE_PORT", 9900))


	c.Loglevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8800"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
