package controllers

import (
	"net/http"
	"time"
	"tower_troops/models"

	"github.com/gin-gonic/gin"
)

var tower *models.Tower

func SetupRoute(r *gin.Engine) *gin.Engine {
  routerGroup := r.Group("/")

  routerGroup.GET("/ping", PingHandler)

  routerGroup.GET("/towers/:id", TowersShowHandler)
  routerGroup.POST("/towers", TowersCreationHandler)
  routerGroup.PUT("/towers/upgrade-armor", TowersUpgradeArmorHandler)
  routerGroup.PUT("/towers/upgrade-hp", TowersUpgradeHPHandler)

  return r
}

func PingHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"msg": "OK", "timestamp": time.Now().Unix()})
}

func TowersCreationHandler(c *gin.Context) {
  tower = models.CreateDefaultTower()

  c.JSON(http.StatusCreated, tower.Serialize())
}

func TowersShowHandler(c *gin.Context) {
  tower = models.CreateDefaultTower()

  c.JSON(http.StatusOK, tower.Serialize())
}

func TowersUpgradeArmorHandler(c *gin.Context) {
  if (tower == nil) {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Tower doesn't exist"})
    return
  }

  if (tower.UpgradeArmor()) {
    c.JSON(http.StatusOK, tower.Serialize())
  } else {
    c.JSON(http.StatusNotAcceptable, gin.H{"error": "Upgrade armor failed"})
  }
}

func TowersUpgradeHPHandler(c *gin.Context) {
  if (tower == nil) {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Tower doesn't exist"})
    return
  }

  if (tower.UpgradeHP()) {
    c.JSON(http.StatusOK, tower.Serialize())
  } else {
    c.JSON(http.StatusNotAcceptable, gin.H{"error": "Upgrade HP failed"})
  }
}
