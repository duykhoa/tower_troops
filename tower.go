package main

type Tower struct {
  HPLevel int
  ArmorLevel int
  Troops []Troop
  Golds int
  Missles []Missle
  UpgradeHPScheme  map[int]UpgradeValue
  UpgradeArmorScheme  map[int]UpgradeValue
}

func CreateDefaultTower() *Tower {
  return &Tower {
    HPLevel: 0,
    ArmorLevel: 0,
    Troops: make([]Troop, 0),
    Golds: 100,
    Missles: make([]Missle, 0),
    UpgradeHPScheme: map[int]UpgradeValue {
      0: { Value: 100, Cost: 0 },
      1: { Value: 150, Cost: 100 },
      2: { Value: 200, Cost: 200 },
      3: { Value: 250, Cost: 300 },
    },
    UpgradeArmorScheme: map[int]UpgradeValue {
      0: { Value: 5, Cost: 0 },
      1: { Value: 7, Cost: 100 },
      2: { Value: 9, Cost: 100 },
      3: { Value: 11, Cost: 200 },
    },
  }
}

func (tower *Tower) MaxArmorLevel() int {
  return len(tower.UpgradeArmorScheme)
}

func (tower *Tower) MaxHPLevel() int {
  return len(tower.UpgradeArmorScheme)
}

func (tower *Tower) UpgradeArmor() bool {
  upgradeLevel := tower.ArmorLevel + 1

  if (upgradeLevel < tower.MaxArmorLevel()) {
    var Value = tower.UpgradeArmorScheme[upgradeLevel]

    if (Value.Cost <= tower.Golds) {
      tower.Golds -= Value.Cost
      tower.ArmorLevel = upgradeLevel

      return true
    }
  }

  return false
}

func (tower *Tower) UpgradeHP() bool {
  upgradeLevel := tower.HPLevel + 1

  if (upgradeLevel < tower.MaxHPLevel()) {
    var Value = tower.UpgradeHPScheme[upgradeLevel]

    if (Value.Cost <= tower.Golds) {
      tower.Golds -= Value.Cost
      tower.HPLevel = upgradeLevel

      return true
    }
  }

  return false
}

func (tower *Tower) HP() int {
  return tower.UpgradeHPScheme[tower.HPLevel].Value
}

func (tower *Tower) Armor() int {
  return tower.UpgradeArmorScheme[tower.ArmorLevel].Value
}

func(tower *Tower) BuyMissle(missle *Missle) {
  if (missle.Cost <= tower.Golds) {
    cloneTower := *missle.Clone()
    cloneTower.SetTower(tower)

    tower.Missles = append(tower.Missles, cloneTower)
    tower.Golds -= missle.Cost
  }
}
