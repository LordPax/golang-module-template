package core

import (
	"fmt"

	"gorm.io/gorm"
)

type IModel[T any] interface {
	FindByID(id string, fields []string) (T, error)
	FindOneBy(field string, value any, fields []string) (T, error)
	Create(entity T) error
	DeleteByID(id string) error
	DeleteBy(field string, value any) error
	UpdateByID(id string, updates T) error
	CountBy(field string, value any) (int64, error)
	GetDB() *gorm.DB
	SetDB(db *gorm.DB)
	Migrate() error
}

type Model[T any] struct {
	*Provider
	db *gorm.DB
}

func NewModel[T any](name string) *Model[T] {
	return &Model[T]{
		Provider: NewProvider(name),
	}
}

func (m *Model[T]) Migrate() error {
	fmt.Printf("Migrating %s\n", m.GetName())
	return m.db.AutoMigrate(new(T))
}

func (m *Model[T]) GetDB() *gorm.DB {
	return m.db
}

func (m *Model[T]) SetDB(db *gorm.DB) {
	m.db = db
}

func (m *Model[T]) FindByID(id string) (T, error) {
	var item T
	if err := m.db.Model(new(T)).First(&item, "id = ?", id).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (m *Model[T]) FindOneBy(field string, value any) (T, error) {
	var item T
	if err := m.db.Model(new(T)).Where(field, value).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (m *Model[T]) Create(entity T) error {
	return m.db.Model(new(T)).Create(entity).Error
}

func (m *Model[T]) DeleteByID(id string) error {
	return m.db.Model(new(T)).Delete(new(T), "id = ?", id).Error
}

func (m *Model[T]) DeleteBy(field string, value any) error {
	return m.db.Model(new(T)).Where(field, value).Delete(new(T)).Error
}

func (m *Model[T]) UpdateByID(id string, updates any) error {
	return m.db.Model(new(T)).Where("id = ?", id).Updates(updates).Error
}

func (m *Model[T]) Save(entity T) error {
	return m.db.Model(new(T)).Save(entity).Error
}

func (m *Model[T]) CountBy(field string, value any) (int64, error) {
	var count int64
	if err := m.db.Model(new(T)).Where(field, value).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
