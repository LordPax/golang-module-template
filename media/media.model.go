package media

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
)

type MediaModel struct {
	*core.Model[*Media]
	databaseService *database.DatabaseService
}

func NewMediaModel(module *MediaModule) *MediaModel {
	return &MediaModel{
		Model:           core.NewModel[*Media]("MediaModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
	}
}

func (um *MediaModel) OnInit() error {
	fmt.Printf("Initializing %s\n", um.GetName())
	um.SetDB(um.databaseService.GetDB())
	return um.Migrate()
}

func (um *MediaModel) FindAll(query query.QueryFilter) ([]*Media, error) {
	var items []*Media

	tx := um.databaseService.GetDB().Model(&Media{}).
		Offset(query.GetSkip()).
		Where(query.GetWhere()).
		Order(query.GetSort())

	if query.GetLimit() != 0 {
		tx.Limit(query.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}
