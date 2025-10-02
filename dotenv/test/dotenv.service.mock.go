package dotenv_test

import (
	"golang-api/core"
)

type DotenvServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewDotenvServiceMock(module core.IModule) *DotenvServiceMock {
	return &DotenvServiceMock{
		Provider: core.NewProvider("DotenvService"),
		Mockable: core.NewMockable(),
	}
}

func (ds *DotenvServiceMock) Load() error {
	ds.MethodCalled("Load")
	ds.CallFunc("Load")
	return nil
}

func (ds *DotenvServiceMock) Get(key string) string {
	ds.MethodCalled("Get", key)
	return ds.CallFunc("Get").(string)
}

func (ds *DotenvServiceMock) Set(key, value string) {
	ds.MethodCalled("Set", key, value)
	ds.CallFunc("Set")
}
