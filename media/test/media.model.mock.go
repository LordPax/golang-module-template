package media_test

import (
	"golang-api/core"
	"golang-api/media"
	"golang-api/query"
)

type MediaModelMock struct {
	*core.ModelMock[*media.Media]
}

func NewMediaModelMock(module *media.MediaModule) *MediaModelMock {
	return &MediaModelMock{
		ModelMock: core.NewModelMock[*media.Media]("MediaModel"),
	}
}

func (um *MediaModelMock) QueryFindAll(q query.QueryFilter) ([]*media.Media, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.CallFunc("QueryFindAll").([]*media.Media), nil
}
