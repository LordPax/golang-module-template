package database_test

import (
	"golang-api/core"
	"golang-api/database"
	dotenv_test "golang-api/dotenv/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewDatabaseModuleTest() *database.DatabaseModule {
	module := &database.DatabaseModule{
		Module: core.NewModule("DatabaseTestModule"),
	}

	module.AddProvider(dotenv_test.NewDotenvServiceMock(module))
	module.AddProvider(database.NewDatabaseService(module))

	return module
}

func TestSetupDatabaseModule(t *testing.T) {
	module := NewDatabaseModuleTest()
	databaseService := module.Get("DatabaseService").(*database.DatabaseService)
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)

	assert.NotNil(t, module, "Database module should be created")
	assert.NotNil(t, databaseService, "DatabaseService should be created")
	assert.NotNil(t, dotenvService, "DotenvService should be created")
}
