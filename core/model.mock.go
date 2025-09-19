package core

import (
	"gorm.io/gorm"
)

// ModelMock is a generic struct that implements the IModel interface for any type T.
type ModelMock[T any] struct {
	*Provider
}

// NewModelMock creates a new instance of ModelMock for the specified type T.
func NewModelMock[T any](name string) *ModelMock[T] {
	return &ModelMock[T]{
		Provider: NewProvider(name),
	}
}

// Migrate performs automatic migration for the model's schema.
func (m *ModelMock[T]) Migrate() error {
	return nil
}

// GetDB returns the current gorm.DB instance.
func (m *ModelMock[T]) GetDB() *gorm.DB {
	return nil
}

// SetDB sets the gorm.DB instance for the model.
func (m *ModelMock[T]) SetDB(db *gorm.DB) {}

// FindByID retrieves a record by its ID.
func (m *ModelMock[T]) FindByID(id string) (T, error) {
	var item T
	return item, nil
}

// FindOneBy retrieves a record by a specified field and value.
func (m *ModelMock[T]) FindOneBy(field string, value any) (T, error) {
	var item T
	return item, nil
}

// Create inserts a new record into the database.
func (m *ModelMock[T]) Create(entity T) error {
	return nil
}

// DeleteByID deletes a record by its ID.
func (m *ModelMock[T]) DeleteByID(id string) error {
	return nil
}

// DeleteBy deletes records matching a specified field and value.
func (m *ModelMock[T]) DeleteBy(field string, value any) error {
	return nil
}

// UpdateByID updates a record by its ID with the provided updates.
func (m *ModelMock[T]) UpdateByID(id string, updates T) error {
	return nil
}

// CountBy counts the number of records matching a specified field and value.
func (m *ModelMock[T]) CountBy(field string, value any) (int64, error) {
	return 0, nil
}

func (m *ModelMock[T]) ClearTable() error {
	return nil
}
