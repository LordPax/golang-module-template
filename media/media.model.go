package media

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
)

type MediaModel struct {
	*core.Model[*Media]
	databaseService *database.DatabaseService
}

func NewMediaModel(module *MediaModule) *MediaModel {
	service := &MediaModel{
		Model:           core.NewModel[*Media]("MediaModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
	}

	module.On("db:migrate", service.Migrate)

	return service
}

func (um *MediaModel) OnInit() error {
	um.SetDB(um.databaseService.GetDB())
	return nil
}

func (um *MediaModel) QueryFindAll(q query.QueryFilter) ([]*Media, error) {
	var items []*Media

	tx := um.databaseService.GetDB().Model(&Media{}).
		Offset(q.GetSkip()).
		Where(q.GetWhere()).
		Order(q.GetSort())

	if q.GetLimit() != 0 {
		tx.Limit(q.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}
