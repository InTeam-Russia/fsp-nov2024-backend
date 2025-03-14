package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel            string
	PostgresUrl         string
	RedisUrl            string
	SessionCookieSecure bool
	SessionCookieDomain string
	AllowOrigin         []string
	AdminUsername       string
	AdminPassword       string
	MockEvents          bool
	MockFilters         bool
}

func LoadConfigFromEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found")
	}

	sessionCookieSecure, err := strconv.ParseBool(os.Getenv("SESSION_COOKIE_SECURE"))
	if err != nil {
		return nil, err
	}

	mockEvents, err := strconv.ParseBool(os.Getenv("MOCK_EVENTS"))
	if err != nil {
		return nil, err
	}

	mockFilters, err := strconv.ParseBool(os.Getenv("MOCK_FILTERS"))
	if err != nil {
		return nil, err
	}

	return &Config{
		LogLevel:            os.Getenv("LOG_LEVEL"),
		PostgresUrl:         os.Getenv("POSTGRES_URL"),
		RedisUrl:            os.Getenv("REDIS_URL"),
		SessionCookieSecure: sessionCookieSecure,
		SessionCookieDomain: os.Getenv("SESSION_COOKIE_DOMAIN"),
		AllowOrigin:         strings.Split(os.Getenv("ALLOW_ORIGIN"), ","),
		AdminUsername:       os.Getenv("ADMIN_USERNAME"),
		AdminPassword:       os.Getenv("ADMIN_PASSWORD"),
		MockEvents:          mockEvents,
		MockFilters:         mockFilters,
	}, nil
}
