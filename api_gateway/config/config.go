package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production
	Loglevel    string
	HTTPPort    string

	CustomerServiceHost string
	CustomerServicePort int
	PostServiceHost     string
	PostServicePort     int
	ReviewServiceHost   string
	ReviewServicePort   int
	CtxTimeout          int
	PostgresHost        string
	PostgresPort        int
	PostgresUser        string
	PostgresPassword    string
	PostgresDB          string
	AuthConfigPath      string
	SignInKey           string
	CsvFilePath         string
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

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_USER_PASSWORD", "123"))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "customerdb"))

	c.Loglevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8800"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_PATH", "./config/auth.conf"))
	c.CsvFilePath = cast.ToString(getOrReturnDefault("CSV_PATH", "./config/auth.csv"))
	c.SignInKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "UmarSecret"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
