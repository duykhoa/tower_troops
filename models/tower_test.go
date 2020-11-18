package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateDefaultTower(t *testing.T) {
  tower := CreateDefaultTower()

  assert.Equal(t, 100, tower.HP())
  assert.Equal(t, 0, tower.HPLevel)
  assert.Empty(t, tower.Troops)
}

func TestUpgradeTowerHP(t *testing.T) {
  tower := CreateDefaultTower()

  result := tower.UpgradeHP()

  assert.True(t, result)
  assert.Equal(t, 150, tower.HP())
  assert.Equal(t, 0, tower.Golds)
}

func TestUpgradeTowerHPWithNotEnoughGold(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 99

  result := tower.UpgradeHP()

  assert.False(t, result)
  assert.Equal(t, 100, tower.HP())
  assert.Equal(t, 99, tower.Golds)
  assert.Equal(t, 0, tower.HPLevel)
}

func TestBuyMissle(t *testing.T) {
  tower := CreateDefaultTower()

  tower.BuyMissle(WoodenArcherMissle)

  assert.Len(t, tower.Missles, 1)
}

func TestBuyMissleWithNotEnoughGold(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 100

  var customMissle = &Missle{
    Cost: 200,
    DamageLevel: 1,
    RangeLevel: 1,
    AttackSpeedLevel: 1,
  }

  tower.BuyMissle(customMissle)

  assert.Empty(t, tower.Missles)
  assert.Equal(t, 100, tower.Golds)
}

func TestUpgradeArmor(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 100

  result := tower.UpgradeArmor()

  assert.True(t, result)
  assert.Equal(t, 1, tower.ArmorLevel)
  assert.Equal(t, 7, tower.Armor())
  assert.Equal(t, 0, tower.Golds)
}

func TestUpgradeArmorWithNotEnoughGold(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 99

  result := tower.UpgradeArmor()

  assert.False(t, result)
  assert.Equal(t, 0, tower.ArmorLevel)
  assert.Equal(t, 5, tower.Armor())
  assert.Equal(t, 99, tower.Golds)
}
