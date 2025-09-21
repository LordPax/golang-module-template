package log

import (
	"fmt"
	"golang-api/core"
	"golang-api/query"
)

type ILogService interface {
	core.IProvider
	FindAll(q query.QueryFilter) ([]*Log, error)
	FindByID(id string) (*Log, error)
	FindOneBy(field string, value any) (*Log, error)
	Create(log *Log) error
	Printf(tags []string, format string, v ...any)
	Errorf(tags []string, format string, v ...any)
}

type LogService struct {
	*core.Provider
	logModel ILogModel
}

func NewLogService(module *LogModule) *LogService {
	return &LogService{
		Provider: core.NewProvider("LogService"),
		logModel: module.Get("LogModel").(ILogModel),
	}
}

func (ls *LogService) FindAll(q query.QueryFilter) ([]*Log, error) {
	return ls.logModel.QueryFindAll(q)
}

func (ls *LogService) FindByID(id string) (*Log, error) {
	return ls.logModel.FindByID(id)
}

func (ls *LogService) FindOneBy(field string, value any) (*Log, error) {
	return ls.logModel.FindOneBy(field, value)
}

func (ls *LogService) Create(log *Log) error {
	return ls.logModel.Create(log)
}

func (ls *LogService) Printf(tags []string, format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	log := NewLog(INFO, tags, text)
	_ = ls.Create(log)
}

func (ls *LogService) Errorf(tags []string, format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	log := NewLog(ERROR, tags, text)
	_ = ls.Create(log)
}
