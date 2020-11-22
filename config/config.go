package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
  PORT string
  DB *gorm.DB
}

const DEFAULT_PORT string = ":8080"
var C *Config

func Load() {
  godotenv.Load()

  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("PG_HOST"),
    os.Getenv("PG_USERNAME"),
    os.Getenv("PG_PASSWORD"),
    os.Getenv("PG_DATABASE"),
    os.Getenv("PG_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if (err != nil) {
    panic("Connect Database failed")
  }

  var port string
  if port = os.Getenv("PORT"); len(port) <= 0 {
    port = DEFAULT_PORT
  }

  C = &Config{
    DB: db,
    PORT: port,
  }
}
