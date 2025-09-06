package token

import (
	"fmt"
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
	return &TokenModel{
		Model:           core.NewModel[*Token]("TokenModel"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
		dotenvService:   module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (um *TokenModel) OnInit() error {
	fmt.Printf("Initializing %s\n", um.GetName())
	um.SetModel(um.databaseService.GetDB().Model(&Token{}))
	return um.Migrate()
}

func (um *TokenModel) DeleteTokensByUserID(userID string) error {
	return um.GetModel().Where("user_id = ?", userID).Delete(&Token{}).Error
}
