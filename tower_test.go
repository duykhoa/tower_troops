package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateDefaultTower(t *testing.T) {
  tower := CreateDefaultTower()

  assert.Equal(t, 100, tower.HP)
  assert.Equal(t, 1, tower.Level)
  assert.Empty(t, tower.Troops)
}

func TestUpgradeTowerHP(t *testing.T) {
  tower := CreateDefaultTower()

  tower.UpgradeHP(50, 70, 1)
  assert.Equal(t, 150, tower.HP)
  assert.Equal(t, 30, tower.Golds)
}

func TestUpgradeTowerHPWithNotEnoughMoney(t *testing.T) {
  tower := CreateDefaultTower()

  tower.UpgradeHP(50, 120, 1)
  assert.Equal(t, 100, tower.HP)
  assert.Equal(t, 100, tower.Golds)
}

func TestUpgradeFollowScheme(t *testing.T) {
  tower := CreateDefaultTower()

  tower.UpdgradeHPScheme()

  assert.Equal(t, 150, tower.HP)
  assert.Equal(t, 30, tower.Golds)
  assert.Equal(t, 1, tower.HPUpgradeLevel)
}

func TestBuyMissle(t *testing.T) {
  tower := CreateDefaultTower()

  tower.BuyMissle(WoodenArcherMissle)

  assert.Len(t, tower.Missles, 1)
}

func TestBuyMissleWithNotEnoughGolds(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 100

  var customMissle = &Missle{
    1, 1, 1, 200,
  }

  tower.BuyMissle(customMissle)

  assert.Len(t, tower.Missles, 0)
  assert.Equal(t, 100, tower.Golds)
}
