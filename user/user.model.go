package user

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"
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
	um.SetModel(um.databaseService.GetDB().Model(&User{}))
	return um.Migrate()
}
