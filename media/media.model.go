package media

import (
	"golang-api/database"
	"golang-api/query"

	"github.com/LordPax/godular/common"
	"github.com/LordPax/godular/core"
)

type IMediaModel interface {
	common.IModel[*Media]
	QueryFindAll(q query.QueryFilter) ([]*Media, error)
}

type MediaModel struct {
	*common.Model[*Media]
	databaseService database.IDatabaseService
}

func NewMediaModel(module core.IModule) *MediaModel {
	service := &MediaModel{
		Model:           common.NewModel[*Media]("MediaModel"),
		databaseService: module.Get("DatabaseService").(database.IDatabaseService),
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
