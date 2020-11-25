package migrations

import (
	"os"
	"testing"
	"tower_troops/config"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
  suite.Suite
}

func (s *Suite) SetupTest() {
  os.Setenv("app_env", "test")
  config.Load()
}

func (s *Suite) TearDownTest() {
  type DBSchemaVersion struct {
    gorm.Model

    Version int
  }

  config.C.DB.Migrator().DropTable(&DBSchemaVersion{})
  os.Unsetenv("app_env")
}

func (s *Suite) TestDBVersionIsZeroForFreshStart() {
  s.Equal(0, GetCurrentDBVersion())
}

func (s *Suite) TestDBVersionAfterRunningMigrate() {
  Migrations = []Migration{
    {
      Version: 1,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
  }

  DBMigrate()

  s.Equal(1, GetCurrentDBVersion())
}

func (s *Suite) TestDBVersionAfterRunningMigrateForAListOfMigrations() {
  Migrations = []Migration{
    {
      Version: 1,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 2,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 10,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
  }

  DBMigrate()

  s.Equal(10, GetCurrentDBVersion())
}

func (s *Suite) TestDBMigrateFromAPreviousDBVersion() {
  Migrations = []Migration{
    {
      Version: 1,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 2,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 4,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 100,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
    },
  }

  config.C.DB.Create(&DBSchemaVersion{ Version: 2 })

  DBMigrate()

  s.Equal(100, GetCurrentDBVersion())
}

func (s *Suite) TestRollBack() {
  Migrations = []Migration{
    {
      Version: 1,
      Migrate: func(db gorm.Migrator) error {
        return nil
      },
      Rollback: func(db gorm.Migrator) error {
        return nil
      },
    },
    {
      Version: 2,
      Migrate: func(db gorm.Migrator) error {
        type DummyTable struct {
          gorm.Model
        }

        return db.CreateTable(&DummyTable{})
      },
      Rollback: func(db gorm.Migrator) error {
        type DummyTable struct {
          gorm.Model
        }

        return db.DropTable(&DummyTable{})
      },
    },
  }

  DBMigrate()

  s.Equal(2, GetCurrentDBVersion())
  s.True(config.C.DB.Migrator().HasTable("dummy_tables"))

  Rollback(1)

  s.Equal(1, GetCurrentDBVersion())
  s.False(config.C.DB.Migrator().HasTable("dummy_tables"))
}

func TestSuite(t *testing.T) {
  _, includeIntegrationTest := os.LookupEnv("RUN_INTEGRATION_TEST")

  if (includeIntegrationTest) {
    suite.Run(t, new(Suite))
  }
}
