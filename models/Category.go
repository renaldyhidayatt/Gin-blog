package models

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ModelCategory struct {
	ID          string         `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" gorm:"unique_index"`
	Slug        string         `json:"slug" gorm:"unique_index"`
	Description string         `json:"description" gorm:"type:text"`
	Article     []ModelArticle `gorm:"many2many:model_articlescategories;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (m *ModelCategory) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now().Local()
	m.Slug = slug.Make(m.Name)
	return nil
}

func (m *ModelCategory) BeforeUpdate(db *gorm.DB) error {
	m.Slug = slug.Make(m.Name)
	m.UpdatedAt = time.Now().Local()
	return nil
}
