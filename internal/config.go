package internal

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	defRefTokExpr = 30 * 24 * time.Hour // default expiration time of refresh token is 1 month
	defIDTokExpr  = 24 * time.Hour      // default expiration time of id token is 1 day
)

type DBConfig struct {
	Username    string
	Password    string
	Host        string
	Port        string
	DBName      string
	MaxAttempts uint8
}

type HTTPServerConfig struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type TokenGenConfig struct {
	Secret      string
	RefreshExpr time.Duration
	IdExpr      time.Duration
	Issuer      string
}

type Config struct {
	DB         DBConfig
	HTTPServer HTTPServerConfig
	TokenGen   TokenGenConfig
}

// Parse panics if required value is not present
func Parse() Config {
	var config Config

	config.DB.Username = mustParse[string]("APP_DB_USERNAME", stringConverter)
	config.DB.Host = parse[string]("APP_DB_HOST", stringConverter, "localhost")
	config.DB.Port = parse[string]("APP_DB_PORT", stringConverter, "5432")
	config.DB.Password = mustParse[string]("APP_DB_PASSWORD", stringConverter)
	config.DB.DBName = mustParse[string]("APP_DB_DATABASE_NAME", stringConverter)
	config.DB.MaxAttempts = parse[uint8]("APP_DB_MAX_CONN_ATTEMPTS", uint8Converter, 5)

	config.HTTPServer.Address = parse[string]("APP_HTTP_ADDRESS", stringConverter, "localhost:8080")
	config.HTTPServer.WriteTimeout = parse[time.Duration]("APP_HTTP_WRITE_TIMEOUT", durationConverter, 0)
	config.HTTPServer.ReadTimeout = parse[time.Duration]("APP_HTTP_READ_TIMEOUT", durationConverter, 0)

	config.TokenGen.Secret = mustParse[string]("APP_TOKEN_SECRET", stringConverter)
	config.TokenGen.Issuer = parse[string]("APP_TOKEN_ISSUER", stringConverter, "test-app")
	config.TokenGen.RefreshExpr = parse[time.Duration]("APP_TOKEN_REFRESH_EXPR", durationConverter, defRefTokExpr)
	config.TokenGen.IdExpr = parse[time.Duration]("APP_TOKEN_ACCESS_EXPR", durationConverter, defIDTokExpr)

	return config
}

func parse[T any](varName string, converter func(val string) (T, error), def T) T {
	valStr, ok := os.LookupEnv(varName)
	if !ok {
		return def
	}

	val, err := converter(valStr)
	if err != nil {
		return def
	}

	return val
}

func mustParse[T any](varName string, converter func(val string) (T, error)) T {
	val, ok := os.LookupEnv(varName)
	if !ok {
		panic(fmt.Sprintf("missing required config parameter: %s", varName))
	}

	retVal, err := converter(val)
	if err != nil {
		panic(fmt.Sprintf(`failed on parsing "%s" due to %v`, varName, err))
	}
	return retVal
}

func stringConverter(val string) (string, error) {
	return val, nil
}

func uint8Converter(val string) (uint8, error) {
	res, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(res), err
}

func durationConverter(val string) (time.Duration, error) {
	return time.ParseDuration(val)
}
