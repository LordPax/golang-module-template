package log

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
)

type LogModel struct {
	*core.Model[*Log]
	databaseService *database.DatabaseService
}

func NewLogModel(module *LogModule) *LogModel {
	return &LogModel{
		Model:           core.NewModel[*Log]("LogModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
	}
}

func (um *LogModel) OnInit() error {
	um.SetDB(um.databaseService.GetDB())
	return um.Migrate()
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
