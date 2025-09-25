package database_test

import (
	"golang-api/core"
	"golang-api/database"

	"gorm.io/gorm"
)

type DatabaseServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewDatabaseServiceMock(module *database.DatabaseModule) *DatabaseServiceMock {
	return &DatabaseServiceMock{
		Provider: core.NewProvider("DatabaseService"),
		Mockable: core.NewMockable(),
	}
}

func (ds *DatabaseServiceMock) Connect() error {
	ds.MethodCalled("Connect")
	ds.CallFunc("Connect")
	return nil
}

func (ds *DatabaseServiceMock) GetDB() *gorm.DB {
	return nil
}

func (ds *DatabaseServiceMock) Close() error {
	ds.MethodCalled("Close")
	ds.CallFunc("Close")
	return nil
}

func (ds *DatabaseServiceMock) Migrate(models ...any) error {
	ds.MethodCalled("Migrate", models)
	ds.CallFunc("Migrate")
	return nil
}

func (ds *DatabaseServiceMock) Table(name string) *gorm.DB {
	return nil
}
