package core

// import "gorm.io/gorm"

// type IEntity interface {
// 	GetModel() *gorm.DB
// 	SetModel(db *gorm.DB)
// 	Save() error
// }

// type Entity struct {
// 	model *gorm.DB
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
// 	err := e.model.Save(e).Error
// 	return err
// }
