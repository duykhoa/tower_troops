package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpgradeRange(t *testing.T) {
  tower := CreateDefaultTower()
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeRange()

  assert.True(t, result)
  assert.Equal(t, 1, missle.RangeLevel)
}

func TestUpgradeRangeWithNotEnoughGolds(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 69
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeRange()

  assert.False(t, result)
  assert.Equal(t, 0, missle.RangeLevel)
}

func TestUpgradeRange_WhenNoTowerAttachedTo(t *testing.T) {
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate)
  result := missle.UpgradeRange()

  assert.False(t, result)
  assert.Equal(t, 0, missle.RangeLevel)
}

func TestUpgradeDamage(t *testing.T) {
  tower := CreateDefaultTower()
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeDamage()

  assert.True(t, result)
  assert.Equal(t, 1, missle.DamageLevel)
}

func TestUpgradeDamageWithNotEnoughGolds(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 69
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeDamage()

  assert.False(t, result)
  assert.Equal(t, 0, missle.DamageLevel)
}

func TestUpgradeDamage_WhenNoTowerAttachedTo(t *testing.T) {
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate)

  result := missle.UpgradeDamage()

  assert.False(t, result)
  assert.Equal(t, 0, missle.DamageLevel)
}

func TestUpgradeAttackSpeed(t *testing.T) {
  tower := CreateDefaultTower()
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeAttackSpeed()

  assert.True(t, result)
  assert.Equal(t, 1, missle.AttackSpeedLevel)
}

func TestUpgradeAttackSpeedWithNotEnoughGolds(t *testing.T) {
  tower := CreateDefaultTower()
  tower.Golds = 49
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate).SetTower(tower)

  result := missle.UpgradeAttackSpeed()

  assert.False(t, result)
  assert.Equal(t, 0, missle.AttackSpeedLevel)
}

func TestUpgradeAttackSpeed_WhenNoTowerAttachedTo(t *testing.T) {
  missle := CreateMissleFromTemplate(WoodenArcherMissleTemplate)

  result := missle.UpgradeAttackSpeed()

  assert.False(t, result)
  assert.Equal(t, 0, missle.AttackSpeedLevel)
}

