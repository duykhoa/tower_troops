package controllers

import (
	"net/http"
	"strings"
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
  routerGroup.POST("/towers/missles", TowersBuyMissleHPHandler)

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
  if (tower == nil) {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Tower doesn't exist"})
    return
  }

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

type BuyMissleJsonRequest struct {
  Name string `json:"missleName"`
}

func TowersBuyMissleHPHandler(c *gin.Context) {
  if (tower == nil) {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Tower doesn't exist"})
    return
  }

  var req BuyMissleJsonRequest

  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  switch strings.ToLower(req.Name) {
  case "wood":
    tower.BuyMissle(models.WoodenArcherMissle)
  default:
    c.JSON(http.StatusBadRequest, gin.H{"error": "No error with name " + req.Name})
  }

  c.JSON(http.StatusOK, tower.Serialize())
}
