package media_test

import (
	"golang-api/core"
	"golang-api/media"
	"golang-api/query"
	"io"
)

type MediaServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewMediaServiceMock(module core.IModule) *MediaServiceMock {
	return &MediaServiceMock{
		Provider: core.NewProvider("MediaService"),
		Mockable: core.NewMockable(),
	}
}

func (us *MediaServiceMock) FindAll(q query.QueryFilter) ([]*media.Media, error) {
	us.MethodCalled("FindAll", q)
	return us.CallFunc("FindAll").([]*media.Media), nil
}

func (us *MediaServiceMock) FindByID(id string) (*media.Media, error) {
	us.MethodCalled("FindByID", id)
	return us.CallFunc("FindByID").(*media.Media), nil
}

func (us *MediaServiceMock) FindOneBy(field string, value any) (*media.Media, error) {
	us.MethodCalled("FindOneBy", field, value)
	return us.CallFunc("FindOneBy").(*media.Media), nil
}

func (us *MediaServiceMock) Create(media *media.Media) error {
	us.MethodCalled("Create", media)
	us.CallFunc("Create")
	return nil
}

func (cs *MediaServiceMock) UploadMedia(
	file io.Reader,
	fileName string,
	fileType string,
	fileSize int64,
	containerName string,
) (*media.Media, error) {
	cs.MethodCalled("UploadMedia", file, fileName, fileType, fileSize, containerName)
	return cs.CallFunc("UploadMedia").(*media.Media), nil
}
