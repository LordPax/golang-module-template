package log

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
)

type ILogModel interface {
	core.IModel[*Log]
	QueryFindAll(q query.QueryFilter) ([]*Log, error)
}

type LogModel struct {
	*core.Model[*Log]
	databaseService database.IDatabaseService
}

func NewLogModel(module core.IModule) *LogModel {
	service := &LogModel{
		Model:           core.NewModel[*Log]("LogModel"),
		databaseService: module.Get("DatabaseService").(database.IDatabaseService),
	}

	module.On("db:migrate", service.Migrate)

	return service
}

func (um *LogModel) OnInit() error {
	um.SetDB(um.databaseService.GetDB())
	return nil
}

func (um *LogModel) QueryFindAll(q query.QueryFilter) ([]*Log, error) {
	var items []*Log

	tx := um.databaseService.GetDB().Model(&Log{}).
		Offset(q.GetSkip()).
		Where(q.GetWhere()).
		Order(q.GetSort())

	if q.GetLimit() != 0 {
		tx.Limit(q.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}
