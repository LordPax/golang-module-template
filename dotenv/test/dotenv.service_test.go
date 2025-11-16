package dotenv_test

import (
	"golang-api/dotenv"
	"testing"

	"github.com/LordPax/godular/core"
	"github.com/stretchr/testify/assert"
)

func NewDotenvModuleTest() *dotenv.DotenvModule {
	module := &dotenv.DotenvModule{
		Module: core.NewModule("DotenvTestModule"),
	}

	module.AddProvider(dotenv.NewDotenvServiceWithPath(module, "env.test"))

	return module
}

func TestSetupDotenvModule(t *testing.T) {
	module := NewDotenvModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv.DotenvService)

	assert.NotNil(t, module, "Dotenv module should be created")
	assert.NotNil(t, dotenvService, "DotenvService should be created")
}

func TestDotenvService_GetSet(t *testing.T) {
	module := NewDotenvModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv.DotenvService)

	expectedValue := "TEST_VALUE"
	expectedKey := "TEST_KEY"
	nonExistentKey := "NON_EXISTENT_KEY"

	dotenvService.Set(expectedKey, expectedValue)
	value := dotenvService.Get(expectedKey)

	if !assert.Equal(t, expectedValue, value, "Value for %s should be %s", expectedKey, expectedValue) {
		return
	}

	nonExistentValue := dotenvService.Get(nonExistentKey)

	if !assert.Equal(t, "", nonExistentValue, "Value for %s should be empty", nonExistentKey) {
		return
	}
}

func TestDotenvService_Load(t *testing.T) {
	module := NewDotenvModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv.DotenvService)

	err := dotenvService.Load()
	if !assert.Nil(t, err, "Load should not return an error") {
		return
	}

	expectedValues := map[string]string{
		"DB_HOST":     "localhost",
		"DB_USER":     "testuser",
		"DB_PASSWORD": "testpassword",
		"DB_NAME":     "testdb",
		"DB_PORT":     "5432",
	}

	for key, expect := range expectedValues {
		value := dotenvService.Get(key)
		if !assert.Equal(t, expect, value, "Value for %s should be %s", key, expect) {
			return
		}
	}
}
