package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelComment struct {
	ID        string       `json:"id" gorm:"primary_key"`
	Content   string       `gorm:"size:2048"`
	Article   ModelArticle `gorm:"foreignKey:ArticleId"`
	ArticleId string
	User      ModelUser `gorm:"foreignKey:UserId"`
	UserId    string

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelComment) BeforeCreate() error {
	m.ID = uuid.NewString()

	m.CreatedAt = time.Now().Local()

	return nil
}

func (m *ModelComment) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()

	return nil
}
