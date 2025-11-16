package log_test

import (
	"golang-api/log"
	"golang-api/query"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type LogModelMock struct {
	*common.ModelMock[*log.Log]
}

func NewLogModelMock(module core.IModule) *LogModelMock {
	return &LogModelMock{
		ModelMock: common.NewModelMock[*log.Log]("LogModel"),
	}
}

func (um *LogModelMock) QueryFindAll(q query.QueryFilter) ([]*log.Log, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.CallFunc("QueryFindAll").([]*log.Log), nil
}
