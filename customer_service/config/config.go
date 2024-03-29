package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	Loglevel          string
	RPCPort           string
	ReviewServiceHost string
	ReviewServicePort int
	PostServiceHost   string
	PostServicePort   int
}

// Load loads environmen vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "database"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5436))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "customerdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "123"))
	
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 9090))

	c.Loglevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9000"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
