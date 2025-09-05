package core

import (
	"fmt"

	"gorm.io/gorm"
)

type IModel[T IEntity] interface {
	FindAll() ([]T, error)
	FindByID(id string) (T, error)
	FindOneBy(field string, value any) (T, error)
	Create(entity T) error
	GetModel() *gorm.DB
	SetModel(db *gorm.DB)
	Migrate() error
}

type Model[T IEntity] struct {
	*Provider
	model *gorm.DB
}

func NewModel[T IEntity](name string) *Model[T] {
	return &Model[T]{
		Provider: NewProvider(name),
	}
}

func (m *Model[T]) Migrate() error {
	fmt.Printf("Migrating %s\n", m.GetName())
	return m.model.AutoMigrate(new(T))
}

func (m *Model[T]) GetModel() *gorm.DB {
	return m.model
}

func (m *Model[T]) SetModel(db *gorm.DB) {
	m.model = db
}

func (m *Model[T]) FindAll() ([]T, error) {
	var items []T
	if err := m.model.Find(&items).Error; err != nil {
		return nil, err
	}
	for _, item := range items {
		item.SetModel(m.model)
	}
	return items, nil
}

func (m *Model[T]) FindByID(id string) (T, error) {
	var item T
	if err := m.model.First(&item, "id = ?", id).Error; err != nil {
		return item, err
	}
	item.SetModel(m.model)
	return item, nil
}

func (m *Model[T]) FindOneBy(field string, value any) (T, error) {
	var item T
	if err := m.model.Where(field, value).First(&item).Error; err != nil {
		return item, err
	}
	item.SetModel(m.model)
	return item, nil
}

func (m *Model[T]) Create(entity T) error {
	return m.model.Create(entity).Error
}
