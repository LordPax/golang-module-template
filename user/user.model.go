package user

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"

	"gorm.io/gorm"
)

type UserModel struct {
	*core.Provider
	databaseService *database.DatabaseService
	model           *gorm.DB
}

func NewUserModel(dbService *database.DatabaseService) *UserModel {
	return &UserModel{
		Provider:        core.NewProvider("UserModel"),
		databaseService: dbService,
	}
}

func (um *UserModel) OnInit() error {
	fmt.Printf("Initializing %s\n", um.GetName())
	um.model = um.databaseService.GetDB().Model(&User{})
	return um.migrate()
}

func (um *UserModel) migrate() error {
	fmt.Printf("Migrating %s\n", um.GetName())
	return um.databaseService.Migrate(&User{})
}

func (um *UserModel) FindAll() ([]*User, error) {
	var users []*User
	err := um.model.Find(&users).Error
	for _, user := range users {
		user.model = um.model
	}
	return users, err
}

func (um *UserModel) FindByID(id string) (*User, error) {
	var user User
	err := um.model.First(&user, "id = ?", id).Error
	user.model = um.model
	return &user, err
}

func (um *UserModel) FindOneBy(field string, value any) (*User, error) {
	var user User
	err := um.model.Where(field, value).First(&user).Error
	user.model = um.model
	return &user, err
}

func (um *UserModel) Create(user *User) error {
	return um.model.Create(user).Error
}
