package log_test

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/query"
)

type LogModelMock struct {
	*core.ModelMock[*log.Log]
}

func NewLogModelMock(module core.IModule) *LogModelMock {
	return &LogModelMock{
		ModelMock: core.NewModelMock[*log.Log]("LogModel"),
	}
}

func (um *LogModelMock) QueryFindAll(q query.QueryFilter) ([]*log.Log, error) {
	um.MethodCalled("QueryFindAll", q)
	return um.CallFunc("QueryFindAll").([]*log.Log), nil
}
