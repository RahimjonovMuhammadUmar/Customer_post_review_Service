package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string // develop, staging, production
	Loglevel          string
	HTTPPort           string
	
	ProductServiceHost  string
	ProductServicePort  int
	CtxTimeout int
}

// Load loads environmen vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.ProductServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.ProductServicePort = cast.ToInt(getOrReturnDefault("PRODUCT_SERVICE_PORT", 9000))

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
