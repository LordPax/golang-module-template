package token

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
)

type ITokenModel interface {
	core.IModel[*Token]
	DeleteByUserID(userID string) error
}

type TokenModel struct {
	*core.Model[*Token]
	databaseService database.IDatabaseService
	dotenvService   dotenv.IDotenvService
}

func NewTokenModel(module core.IModule) *TokenModel {
	service := &TokenModel{
		Model:           core.NewModel[*Token]("TokenModel"),
		databaseService: module.Get("DatabaseService").(database.IDatabaseService),
		dotenvService:   module.Get("DotenvService").(dotenv.IDotenvService),
	}

	module.On("db:migrate", service.Migrate)

	return service
}

func (um *TokenModel) OnInit() error {
	um.SetDB(um.databaseService.GetDB())
	return nil
}

func (um *TokenModel) DeleteByUserID(userID string) error {
	return um.GetDB().Model(&Token{}).Where("user_id = ?", userID).Delete(&Token{}).Error
}
