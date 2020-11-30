package models

import (
	"os"
	"testing"
	"tower_troops/config"
	"tower_troops/migrations"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TowerRepositorySuite struct {
  suite.Suite
}

type DBSchemaVersion struct {
  gorm.Model

  Version int
}

var _DB *gorm.DB

func (s *TowerRepositorySuite) SetupTest() {
  os.Setenv("app_env", "test")
  config.Load()
  migrations.DBMigrate()

  _DB = config.C.DB
  config.C.DB = config.C.DB.Begin()
}

func (s *TowerRepositorySuite) TearDownTest() {
  config.C.DB.Rollback()
  config.C.DB = _DB
  config.C.DB.Delete(&DBSchemaVersion{})
  os.Unsetenv("app_env")
}

func (s *TowerRepositorySuite) TestCreateTower() {
  user := &User {
    Name: "Kevin",
    UserName: "duykhoa",
    EncryptedPassword: "password",
  }

  config.C.DB.Create(user)

  tower := CreateDefaultTower()
  tower.User = user

  result := config.C.DB.Create(tower)

  var users []User

  config.C.DB.Find(&users)

  s.Len(users, 1)
  s.NoError(result.Error)
}

func TestSuite(t *testing.T) {
  suite.Run(t, new(TowerRepositorySuite))
}
