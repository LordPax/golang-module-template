package core

// import "gorm.io/gorm"

// type IEntity interface {
// 	GetModel() *gorm.DB
// 	SetModel(db *gorm.DB)
// 	Save() error
// }

// type Entity struct {
// 	model *gorm.DB `json:"-" gorm:"-"`
// }

// func NewEntity(db *gorm.DB) *Entity {
// 	return &Entity{
// 		model: db,
// 	}
// }

// func (e *Entity) GetModel() *gorm.DB {
// 	return e.model
// }

// func (e *Entity) SetModel(db *gorm.DB) {
// 	e.model = db
// }

// func (e *Entity) Save() error {
// 	return e.model.Save(e).Error
// }
