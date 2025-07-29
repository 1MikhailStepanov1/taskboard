package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type AppConfig struct {
	SwaggerPath         string
	HTTPPort            int
	HTTPShutdownTimeout time.Duration
	LogLevel            string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type SecurityConfig struct {
	JWTSecret   string
	JWTDuration time.Duration
}

type Config struct {
	App      AppConfig
	DB       DBConfig
	Security SecurityConfig
}

func New() *Config {
	return &Config{
		App: AppConfig{
			SwaggerPath:         getEnv("SWAGGER_PATH", "/app/http/swagger.json"),
			HTTPPort:            getEnvAsInt("HTTP_APP_PORT", 6666),
			HTTPShutdownTimeout: time.Duration(getEnvAsInt("HTTP_SHUTDOWN_TIMEOUT", 10)),
			LogLevel:            getEnv("LOG_LEVEL", "PROD"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Database: getEnv("DB_DATABASE", "postgres"),
		},
		Security: SecurityConfig{
			JWTSecret:   getEnv("JWT_SECRET", "pipipupu"),
			JWTDuration: time.Duration(getEnvAsInt("JWT_DURATION", 60)),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
