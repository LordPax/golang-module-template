package media_test

import (
	"golang-api/media"
	"golang-api/query"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type MediaModelMock struct {
	*common.ModelMock[*media.Media]
}

func NewMediaModelMock(module core.IModule) *MediaModelMock {
	return &MediaModelMock{
		ModelMock: common.NewModelMock[*media.Media]("MediaModel"),
	}
}

func (um *MediaModelMock) QueryFindAll(q query.QueryFilter) ([]*media.Media, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.CallFunc("QueryFindAll").([]*media.Media), nil
}
