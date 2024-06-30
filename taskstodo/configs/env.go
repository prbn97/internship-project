package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// create a type to hold the enviroment variables
type Config struct {
	PublicHost string
	Port       string

	DBuser                 string
	DBpassWord             string
	DBaddress              string
	DBname                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

// create a object with the enviroment variables
func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBuser:                 getEnv("DB_USER", "root"),
		DBpassWord:             getEnv("DB_PASSWORD", "mypassword"),
		DBaddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBname:                 getEnv("DB_NAME", "internship_project"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

// hold this object to use in cmd/ main.go
var Envs = initConfig()

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
