package media

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(255)"`
	Url       string     `json:"url" gorm:"type:text;not null"`
	Type      string     `json:"type" gorm:"type:varchar(50)"`
	Size      int64      `json:"size" gorm:"type:bigint"`
	Container string     `json:"container" gorm:"type:varchar(100)"`
	UserID    string     `json:"user_id" gorm:"type:varchar(100)"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (m *Media) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
