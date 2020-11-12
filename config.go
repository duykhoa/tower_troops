package main

type TowerHPUpgradeSetting struct {
  HPAmount int
  cost int
}

var TowerHPUpgradeScheme = map[int]TowerHPUpgradeSetting {
  1: { HPAmount: 50, cost: 70 },
  2: { HPAmount: 50, cost: 70 },
  3: { HPAmount: 50, cost: 80 },
}

var MaxTowerHPUpdateLevel = len(TowerHPUpgradeScheme)
