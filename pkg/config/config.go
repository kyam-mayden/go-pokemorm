package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get(key string, backup string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  envValue := os.Getenv(key)

  if len(envValue) > 0 {
    return envValue
  }

  return backup
}

// maybe a config struct for the whole app
