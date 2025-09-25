package gin_test

import (
	"golang-api/core"
	ginM "golang-api/gin"

	"github.com/gin-gonic/gin"
)

type GinServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewGinServiceMock(module *ginM.GinModule) *GinServiceMock {
	return &GinServiceMock{
		Provider: core.NewProvider("GinService"),
		Mockable: core.NewMockable(),
	}
}

func (gs *GinServiceMock) GetGroup() *gin.RouterGroup {
	return nil
}

func (gs *GinServiceMock) Cors() gin.HandlerFunc {
	gs.MethodCalled("Cors")
	gs.CallFunc("Cors")
	return nil
}

func (gs *GinServiceMock) Swagger() {
	gs.MethodCalled("Swagger")
	gs.CallFunc("Swagger")
}

func (gs *GinServiceMock) InitEngine() {
	gs.MethodCalled("InitEngine")
	gs.CallFunc("InitEngine")
}

func (gs *GinServiceMock) Run() {
	gs.MethodCalled("Run")
	gs.CallFunc("Run")
}
