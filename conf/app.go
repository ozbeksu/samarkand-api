package conf

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type AppConfig struct {
	env   string
	port  int
	debug bool
	cors  cors.Config
}

func NewAppConfig() *AppConfig {
	env := os.Getenv("APP_ENV")
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	return &AppConfig{
		env:   env,
		port:  port,
		debug: isDebugging(),
		cors: cors.Config{
			AllowOrigins: os.Getenv("CORS_ORIGIN"),
			AllowHeaders: os.Getenv("CORS_HEADERS"),
		},
	}
}

func (c AppConfig) Env() string {
	return c.env
}

func (c AppConfig) IsProd() bool {
	return c.env == "production"
}

func (c AppConfig) IsDev() bool {
	return c.env == "dev"
}

func (c AppConfig) Debug() bool {
	return c.debug
}

func (c AppConfig) Port() string {
	return fmt.Sprintf(":%d", c.port)
}

func (c AppConfig) Cors() cors.Config {
	return c.cors
}

func LoadEnvConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	f, _ := os.Getwd()
	envFile := fmt.Sprintf("%s/%s/.env.%s", filepath.Dir(f), filepath.Base(f), env)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env.dev file")
	}
}

func isDebugging() bool {
	return os.Getenv("APP_DEBUG") == "true"
}
