package user

import (
	"golang-api/core"
	"golang-api/database"
	"golang-api/query"
	"golang-api/utils"
	"time"
)

type IUserModel interface {
	core.IModel[*User]
	QueryFindAll(q query.QueryFilter) ([]*User, error)
	DeleteByID(id string) error
	CountAll() (int64, error)
}

type UserModel struct {
	*core.Model[*User]
	databaseService database.IDatabaseService
}

func NewUserModel(module *UserModule) *UserModel {
	service := &UserModel{
		Model:           core.NewModel[*User]("UserModel"),
		databaseService: module.Get("DatabaseService").(database.IDatabaseService),
	}

	module.On("db:migrate", service.Migrate)

	return service
}

func (um *UserModel) OnInit() error {
	um.SetDB(um.databaseService.GetDB())
	return nil
}

func (um *UserModel) QueryFindAll(q query.QueryFilter) ([]*User, error) {
	var items []*User

	tx := um.databaseService.GetDB().Model(&User{}).
		Where("deleted_at IS NULL").
		Select(secureFields).
		Offset(q.GetSkip()).
		Where(q.GetWhere()).
		Order(q.GetSort())

	if q.GetLimit() != 0 {
		tx.Limit(q.GetLimit())
	}

	err := tx.Find(&items).Error
	return items, err
}

func (um *UserModel) DeleteByID(id string) error {
	deletedAt := time.Now()
	user := User{
		DeletedAt: &deletedAt,
		Username:  "[deleted]",
		Firstname: "[deleted]",
		Lastname:  "[deleted]",
		Email:     "[deleted]",
		Profile:   "https://api.dicebear.com/9.x/initials/svg?seed=deleted",
	}
	user.HashPassword(utils.GenerateString(12))
	return um.UpdateByID(id, &user)
}

func (um *UserModel) CountAll() (int64, error) {
	var count int64

	err := um.databaseService.GetDB().Model(&User{}).
		Where("deleted_at IS NULL").
		Count(&count).Error

	return count, err
}
