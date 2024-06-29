package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// create a type to hold the enviroment variables

type Config struct {
	PublicHost string
	Port       string

	DBuser     string
	DBpassWord string
	DBaddress  string
	DBname     string
}

// and the create a object with the enviroment variables
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBuser:     getEnv("DB_USER", "root"),
		DBpassWord: getEnv("DB_PASSWORD", "mypassword"),
		DBaddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBname:     getEnv("DB_NAME", "internship_project"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
