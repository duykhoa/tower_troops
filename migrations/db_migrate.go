package migrations

import (
  "errors"
  "tower_troops/config"

  "gorm.io/gorm"
)

type DBSchemaVersion struct {
  gorm.Model

  Version int
}

type Migration struct {
  Version int
  Desc string
  Migrate func(db gorm.Migrator) error
  Rollback func(db gorm.Migrator) error
}

func GetCurrentDBVersion() (currentDBVersionNumber int) {
  migrator := config.C.DB.Migrator()

  if (!migrator.HasTable(&DBSchemaVersion{})) {
    currentDBVersionNumber = 0
  } else {
    var lastestDBSchemaVersion DBSchemaVersion
    result := config.C.DB.Last(&lastestDBSchemaVersion)

    if (!errors.Is(result.Error, gorm.ErrRecordNotFound)) {
      currentDBVersionNumber = lastestDBSchemaVersion.Version
    } else {
      currentDBVersionNumber = 0
    }
  }

  return
}

func DBMigrate() {
  migrator := config.C.DB.Migrator()

  if (!migrator.HasTable(&DBSchemaVersion{})) {
    config.C.DB.Migrator().CreateTable(&DBSchemaVersion{})
  }
  var newerDBVersionNumber int
  var currentDBVersionNumber int = GetCurrentDBVersion()

  for _, migration := range(Migrations) {
    if (migration.Version > currentDBVersionNumber) {
      newerDBVersionNumber = migration.Version
      migration.Migrate(migrator)
    }
  }

  if (newerDBVersionNumber > currentDBVersionNumber) {
    config.C.DB.Create(&DBSchemaVersion{
      Version: newerDBVersionNumber,
    })
  }
}

func Rollback(noSteps int) {
  migrator := config.C.DB.Migrator()
  currentDBVersionNumber := GetCurrentDBVersion()

  for _, migration := range(Migrations) {
    if (migration.Version > currentDBVersionNumber - noSteps && migration.Version <= currentDBVersionNumber) {
      migration.Rollback(migrator)
    }
  }

  if (currentDBVersionNumber - noSteps < currentDBVersionNumber) {
    config.C.DB.Create(&DBSchemaVersion{
      Version: currentDBVersionNumber - noSteps,
    })
  }
}

var Migrations = []Migration {
  {
    Version: 1,
    Migrate: func(db gorm.Migrator) error {
      type User struct {
        gorm.Model

        Name string              `gorm:"not null"`
        UserName string          `gorm:"not null"`
        EncryptedPassword string `gorm:"not null"`
      }

      type Troop struct {
        gorm.Model

        TowerID uint
      }

      type Missle struct {
        gorm.Model

        TowerID uint
        Cost int                 `gorm:"not null"`
        Name string              `gorm:"not null"`
        DamageLevel int          `gorm:"not null, default:0"`
        RangeLevel int           `gorm:"not null, default:0"`
        AttackSpeedLevel int     `gorm:"not null, default:0"`
      }

      type Tower struct {
        gorm.Model

        HPLevel int              `gorm:"not null"`
        ArmorLevel int           `gorm:"not null"`
        Troops []Troop           `gorm:"not null"`
        Golds int                `gorm:"not null"`
        Missles []Missle
      }

      return db.AutoMigrate(&User{}, &Tower{}, &Troop{}, &Missle{})
    },
    Rollback: func(db gorm.Migrator) error {
      type User struct {}
      type Troop struct {}
      type Missle struct {}
      type Tower struct {}

      return db.DropTable(&User{}, &Troop{}, &Missle{}, &Tower{})
    },
  },
}

