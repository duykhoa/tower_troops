package models

import "gorm.io/gorm"

type Troop struct {
  gorm.Model

  TowerID uint
}
