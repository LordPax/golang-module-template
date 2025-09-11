package log

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Type      string     `json:"type" gorm:"type:varchar(100)"`
	Tags      []string   `json:"tags" gorm:"json"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewLog(logType string, tags []string, message string) *Log {
	return &Log{
		ID:      uuid.New().String(),
		Type:    logType,
		Tags:    tags,
		Message: message,
	}
}

func (l *Log) AddTag(tag string) {
	l.Tags = append(l.Tags, tag)
}
