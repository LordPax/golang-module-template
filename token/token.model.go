package token

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
)

type TokenModel struct {
	*core.Model[*Token]
	databaseService *database.DatabaseService
	dotenvService   *dotenv.DotenvService
}

func NewTokenModel(module *TokenModule) *TokenModel {
	service := &TokenModel{
		Model:           core.NewModel[*Token]("TokenModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
		dotenvService:   module.Get("DotenvService").(*dotenv.DotenvService),
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
