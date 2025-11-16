package media_test

import (
	"io"

	"github.com/LordPax/godular/core"
)

type OpenstackServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewOpenstackServiceMock(module core.IModule) *OpenstackServiceMock {
	return &OpenstackServiceMock{
		Provider: core.NewProvider("OpenstackService"),
		Mockable: core.NewMockable(),
	}
}

func (o *OpenstackServiceMock) Authenticate() error {
	o.MethodCalled("Authenticate")
	o.CallFunc("Authenticate")
	return nil
}

func (o *OpenstackServiceMock) CreateContainerIfNotExist(containerName string) error {
	o.MethodCalled("CreateContainerIfNotExist", containerName)
	o.CallFunc("CreateContainerIfNotExist")
	return nil
}

func (o *OpenstackServiceMock) UploadFile(
	file io.Reader,
	objectName string,
	containerName string,
) (string, error) {
	o.MethodCalled("UploadFile", file, objectName, containerName)
	return o.CallFunc("UploadFile").(string), nil
}
