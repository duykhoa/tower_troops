package main

import (
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

func setupRoute(r *gin.Engine) *gin.Engine {
  routerGroup := r.Group("/v1")

  routerGroup.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"msg": "OK", "timestamp": time.Now().Unix()})
  })

  return r
}

func main() {
  //gin.SetMode(gin.ReleaseMode)

  r := gin.Default()
  setupRoute(r).Run(":8080")
}
