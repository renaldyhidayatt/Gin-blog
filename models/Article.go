package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ModelArticle struct {
	ID          string `json:"id" gorm:"primary_key"`
	Slug        string `gorm:"unique_index"`
	Title       string `json:"title" gorm:"type:varchar; not null"`
	Description string `json:"description" gorm:"type:text; not null"`
	Body        string `json:"body" gorm:"type:varchar; not null"`
	User        ModelUser
	UserID      string          `json:"user_id" gorm:"index"`
	Tags        []ModelTag      `gorm:"many2many:articles_tags;"`
	Categories  []ModelCategory `gorm:"many2many:articles_categories;"`
	Comments    []ModelComment  `gorm:"foreignKey:ArticleId"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func (m *ModelArticle) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Slug = slug.Make(m.Title)
	m.CreatedAt = time.Now().Local()

	return nil
}

func (m *ModelArticle) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Slug = slug.Make(m.Title)
	m.UpdatedAt = time.Now().Local()

	return nil
}
