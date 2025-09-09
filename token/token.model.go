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
	um.SetDB(um.databaseService.GetDB())
	// return nil
	return um.Migrate()
}

func (um *TokenModel) DeleteByUserID(userID string) error {
	return um.GetDB().Model(&Token{}).Where("user_id = ?", userID).Delete(&Token{}).Error
}
