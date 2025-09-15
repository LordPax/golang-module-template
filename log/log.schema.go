package log

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Log struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type" gorm:"type:varchar(100)"`
	Tags      []string  `json:"tags" gorm:"json"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLog(logType string, tags []string, message string) *Log {
	return &Log{
		Type:    logType,
		Tags:    tags,
		Message: message,
	}
}

func (l *Log) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New().String()
	return nil
}

func (l *Log) AddTag(tag string) {
	l.Tags = append(l.Tags, tag)
}
