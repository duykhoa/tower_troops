package main

import (
	"github.com/gin-gonic/gin"
	"tower_troops/controllers"
)

const PORT = ":8080"

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	controllers.SetupRoute(r).Run(PORT)
}
