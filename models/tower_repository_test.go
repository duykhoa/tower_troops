package models

import (
	"fmt"
	"os"
	"testing"
	"tower_troops/config"
	"tower_troops/migrations"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
  suite.Suite
}

type DBSchemaVersion struct {
  gorm.Model

  Version int
}

func (s *Suite) SetupTest() {
  fmt.Println("Setup test")
  os.Setenv("app_env", "test")
  config.Load()
  migrations.DBMigrate()
  config.C.DB.Begin()
}

func (s *Suite) TearDownTest() {
  config.C.DB.Rollback()
  config.C.DB.Migrator().DropTable(&DBSchemaVersion{})
  os.Unsetenv("app_env")
}

func (s *Suite) TestCreateTower() {
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
  _, includeIntegrationTest := os.LookupEnv("RUN_INTEGRATION_TEST")

  if (includeIntegrationTest) {
    suite.Run(t, new(Suite))
  }
}
