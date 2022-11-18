package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelTag struct {
	ID          string         `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" gorm:"type:varchar; not null"`
	Slug        string         `json:"slug" gorm:"type:varchar; not null"`
	Description string         `json:"description" gorm:"type:varchar; not null"`
	Article     []ModelArticle `gorm:"many2many:articles_tags;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (m *ModelTag) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().Local()

	return nil
}

func (m *ModelTag) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()
	return nil
}
