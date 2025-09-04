package core

// import (
// 	"gorm.io/gorm"
// )

// type IModel interface {
// 	// FindAll() ([]IEntity, error)
// 	// FindByID(id string) (IEntity, error)
// 	// FindOneBy(field string, value interface{}) (IEntity, error)
// 	// Create(entity IEntity) error
// 	GetModel() *gorm.DB
// 	SetModel(db *gorm.DB)
// }

// type Model struct {
// 	*Provider
// 	table string
// 	model *gorm.DB
// }

// func NewModel(name string, model *gorm.DB) *Model {
// 	return &Model{
// 		Provider: NewProvider(name),
// 		model:    model,
// 	}
// }

// func (m *Model) GetModel() *gorm.DB {
// 	return m.model
// }

// func (m *Model) SetModel(db *gorm.DB) {
// 	m.model = db
// }

// func (m *Model) FindAll() ([]IEntity, error) {
// 	var items []IEntity
// 	err := m.model.Find(&items).Error
// 	for _, item := range items {
// 		item.SetModel(m.model)
// 	}
// 	return items, err
// }

// func (m *Model) FindByID(id string) (IEntity, error) {
// 	var item IEntity
// 	err := m.model.First(&item, "id = ?", id).Error
// 	item.SetModel(m.model)
// 	return item, err
// }

// func (m *Model) FindOneBy(field string, value any) (IEntity, error) {
// 	var item IEntity
// 	err := m.model.Where(field, value).First(&item).Error
// 	return item, err
// }

// func (m *Model) Create(entity IEntity) error {
// 	return m.model.Create(entity).Error
// }
