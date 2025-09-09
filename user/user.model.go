package user

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
)

type UserModel struct {
	*core.Model[*User]
	databaseService *database.DatabaseService
}

func NewUserModel(module *UserModule) *UserModel {
	return &UserModel{
		Model:           core.NewModel[*User]("UserModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
	}
}

func (um *UserModel) OnInit() error {
	fmt.Printf("Initializing %s\n", um.GetName())
	um.SetDB(um.databaseService.GetDB())
	return um.Migrate()
}

func (um *UserModel) FindAll(query query.QueryFilter) ([]*User, error) {
	var items []*User

	tx := um.databaseService.GetDB().Model(&User{}).
		Where("deleted_at IS NULL").
		Offset(query.GetSkip()).
		Where(query.GetWhere()).
		Order(query.GetSort())

	if query.GetLimit() != 0 {
		tx.Limit(query.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}
