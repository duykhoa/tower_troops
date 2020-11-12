package main

type Tower struct {
  Level int
  HP int
  HPUpgradeLevel int
  Troops []Troop
  Golds int
  Missles []Missle
}

func CreateDefaultTower() *Tower {
  return &Tower {
    1, 100, 0, make([]Troop, 0), 100, make([]Missle, 0),
  }
}

func (tower *Tower) UpgradeHP(hpAmount int, cost int, hpUpgradeLevel int) {
  if (cost <= tower.Golds) {
    tower.HP += hpAmount
    tower.Golds -= cost
    tower.HPUpgradeLevel = hpUpgradeLevel
  }
}

func (tower *Tower) UpdgradeHPScheme() {
  hpUpgradeLevel := tower.HPUpgradeLevel

  if (hpUpgradeLevel < MaxTowerHPUpdateLevel) {
    var hpUpgradeSetting = TowerHPUpgradeScheme[hpUpgradeLevel + 1]

    tower.UpgradeHP(hpUpgradeSetting.HPAmount, hpUpgradeSetting.cost, hpUpgradeLevel + 1)
  }
}

func(tower *Tower) BuyMissle(missle *Missle) {
  if (missle.Cost <= tower.Golds) {
    tower.Missles = append(tower.Missles, *missle.Clone())
    tower.Golds -= missle.Cost
  }
}
