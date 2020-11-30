package models

import "gorm.io/gorm"

type Missle struct {
  gorm.Model

  TowerID uint
  Cost int
  Name string
  DamageLevel int
  RangeLevel int
  Tower *Tower
  AttackSpeedLevel int
  MissleTemplate *MissleTemplate `json:"-" gorm:"-"`
}

type MissleTemplate struct {
  Cost int
  Name string
  DamageLevel int
  RangeLevel int
  AttackSpeedLevel int
  UpgradeDamageLevelScheme map[int]UpgradeValue
  UpgradeRangeLevelScheme map[int]UpgradeValue
  UpgradeAttackSpeedLevelScheme map[int]UpgradeValue
}

var WoodenArcherMissleTemplate = &MissleTemplate{
  Cost: 80,
  Name: "WoodenArcherMissle",
  AttackSpeedLevel: 0,
  DamageLevel: 0,
  RangeLevel: 0,
  UpgradeRangeLevelScheme: map[int]UpgradeValue {
    0: { Value: 4, Cost: 0 },
    1: { Value: 5, Cost: 70 },
    2: { Value: 6, Cost: 200 },
    3: { Value: 7, Cost: 300 },
  },
  UpgradeDamageLevelScheme: map[int]UpgradeValue{
    0: { Value: 7, Cost: 0 },
    1: { Value: 10, Cost: 70 },
    2: { Value: 12, Cost: 80 },
    3: { Value: 13, Cost: 80 },
  },
  UpgradeAttackSpeedLevelScheme: map[int]UpgradeValue{
    0: { Value: 4, Cost: 0 },
    1: { Value: 5, Cost: 50 },
    2: { Value: 7, Cost: 200 },
    3: { Value: 10, Cost: 300 },
  },
}

func CreateMissleFromTemplate(missleTemplate *MissleTemplate) *Missle {
  return &Missle{
    Cost: missleTemplate.Cost,
    DamageLevel: missleTemplate.DamageLevel,
    RangeLevel: missleTemplate.RangeLevel,
    AttackSpeedLevel: missleTemplate.AttackSpeedLevel,
    MissleTemplate: missleTemplate,
  }
}

func (missle *Missle) SetTower(tower *Tower) *Missle {
  missle.Tower = tower
  return missle
}

func (missle *Missle) MaxRangeLevel() int {
  return len(missle.MissleTemplate.UpgradeRangeLevelScheme)
}

func (missle *Missle) MaxDamageLevel() int {
  return len(missle.MissleTemplate.UpgradeDamageLevelScheme)
}

func (missle *Missle) MaxAttackSpeedLevel() int {
  return len(missle.MissleTemplate.UpgradeAttackSpeedLevelScheme)
}

func (missle *Missle) Range() int {
  return missle.MissleTemplate.UpgradeRangeLevelScheme[missle.RangeLevel].Value
}

func (missle *Missle) Damage() int {
  return missle.MissleTemplate.UpgradeDamageLevelScheme[missle.DamageLevel].Value
}

func (missle *Missle) AttackSpeed() int {
  return missle.MissleTemplate.UpgradeAttackSpeedLevelScheme[missle.AttackSpeedLevel].Value
}

func (missle *Missle) UpgradeRange() bool {
  upgradeLevel := missle.RangeLevel + 1

  if (upgradeLevel < missle.MaxRangeLevel() && missle.Tower != nil) {
    var Value = missle.MissleTemplate.UpgradeRangeLevelScheme[upgradeLevel]

    if (Value.Cost <= missle.Tower.Golds) {
      missle.Tower.Golds -= Value.Cost
      missle.RangeLevel = upgradeLevel

      return true
    }
  }

  return false
}

func (missle *Missle) UpgradeDamage() bool {
  upgradeLevel := missle.DamageLevel + 1

  if (upgradeLevel < missle.MaxDamageLevel() && missle.Tower != nil) {
    var Value = missle.MissleTemplate.UpgradeDamageLevelScheme[upgradeLevel]

    if (Value.Cost <= missle.Tower.Golds) {
      missle.Tower.Golds -= Value.Cost
      missle.DamageLevel = upgradeLevel

      return true
    }
  }

  return false
}

func (missle *Missle) UpgradeAttackSpeed() bool {
  upgradeLevel := missle.AttackSpeedLevel + 1

  if (upgradeLevel < missle.MaxAttackSpeedLevel() && missle.Tower != nil) {
    var Value = missle.MissleTemplate.UpgradeAttackSpeedLevelScheme[upgradeLevel]

    if (Value.Cost <= missle.Tower.Golds) {
      missle.Tower.Golds -= Value.Cost
      missle.AttackSpeedLevel = upgradeLevel

      return true
    }
  }

  return false
}

func (missle *Missle) Serialize() *MissleSerializer {
  return &MissleSerializer{
    Damage: missle.Damage(),
    DamageLevel: missle.DamageLevel,
    Range: missle.Range(),
    RangeLevel: missle.RangeLevel,
    AttackSpeed: missle.AttackSpeed(),
    AttackSpeedLevel: missle.AttackSpeedLevel,
    Name: missle.Name,
  }
}
