package main

import (
  "os"
  "tower_troops/config"
  "tower_troops/controllers"
  "tower_troops/migrations"

  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

const PORT = ":8080"

type User struct {
  gorm.Model
  ID int `gorm:"primaryKey"`
  email string
}

func main() {
  config.Load()

  if (len(os.Args) > 1) {
    if (os.Args[1] == "migrate") {
      migrations.DBMigrate()
    } else if (os.Args[1] == "rollback") {
      migrations.Rollback(1)
    }
  } else {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    controllers.SetupRoute(r).Run(PORT)
  }
}
