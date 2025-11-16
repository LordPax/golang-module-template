package log

import (
	"golang-api/database"
	"golang-api/query"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type ILogModel interface {
	common.IModel[*Log]
	QueryFindAll(q query.QueryFilter) ([]*Log, error)
}

type LogModel struct {
	*common.Model[*Log]
	databaseService database.IDatabaseService
}

func NewLogModel(module core.IModule) *LogModel {
	service := &LogModel{
		Model:           common.NewModel[*Log]("LogModel"),
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
