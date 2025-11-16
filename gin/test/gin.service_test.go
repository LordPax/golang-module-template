package gin_test

import (
	dotenv_test "golang-api/dotenv/test"
	"golang-api/gin"
	"testing"

	"github.com/LordPax/godular/core"
	"github.com/stretchr/testify/assert"
)

func NewGinModuleTest() *gin.GinModule {
	module := &gin.GinModule{
		Module: core.NewModule("GinTestModule"),
	}

	module.AddProvider(dotenv_test.NewDotenvServiceMock(module))
	module.AddProvider(gin.NewGinService(module))

	return module
}

func TestSetupGinModule(t *testing.T) {
	module := NewGinModuleTest()
	dotenvService := module.Get("DotenvService").(*dotenv_test.DotenvServiceMock)
	ginService := module.Get("GinService").(*gin.GinService)

	assert.NotNil(t, module, "Gin module should be created")
	assert.NotNil(t, dotenvService, "DotenvService should be created")
	assert.NotNil(t, ginService, "GinService should be created")
}
