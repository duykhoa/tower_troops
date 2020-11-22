package models

import "gorm.io/gorm"

type User struct {
  gorm.Model

  Name string               `gorm:"not null"`
  UserName string           `gorm:"not null"`
  EncryptedPassword string  `gorm:"not null"`
}
