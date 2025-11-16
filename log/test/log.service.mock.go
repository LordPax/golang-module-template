package log_test

import (
	"golang-api/log"
	"golang-api/query"

	"github.com/LordPax/godular/core"
)

type LogServiceMock struct {
	*core.Provider
	*core.Mockable
}

func NewLogServiceMock(module core.IModule) *LogServiceMock {
	return &LogServiceMock{
		Provider: core.NewProvider("LogService"),
		Mockable: core.NewMockable(),
	}
}

func (ls *LogServiceMock) FindAll(q query.QueryFilter) ([]*log.Log, error) {
	ls.MethodCalled("FindAll", q)
	return ls.CallFunc("FindAll").([]*log.Log), nil
}

func (ls *LogServiceMock) FindByID(id string) (*log.Log, error) {
	ls.MethodCalled("FindByID", id)
	return ls.CallFunc("FindByID").(*log.Log), nil
}

func (ls *LogServiceMock) FindOneBy(field string, value any) (*log.Log, error) {
	ls.MethodCalled("FindOneBy", field, value)
	return ls.CallFunc("FindOneBy").(*log.Log), nil
}

func (ls *LogServiceMock) Create(log *log.Log) error {
	ls.MethodCalled("Create", log)
	ls.CallFunc("Create")
	return nil
}

func (ls *LogServiceMock) Printf(tags []string, format string, v ...any) {
	ls.MethodCalled("Printf", tags, format, v)
	ls.CallFunc("Printf")
}

func (ls *LogServiceMock) Errorf(tags []string, format string, v ...any) {
	ls.MethodCalled("Errorf", tags, format, v)
	ls.CallFunc("Errorf")
}
