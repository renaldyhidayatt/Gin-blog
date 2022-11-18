package models

import (
	"ginBlog/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelUser struct {
	ID        string `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"type:varchar; not null"`
	FirstName string `json:"firstname" gorm:"type:varchar;not null"`
	LastName  string `json:"lastname" gorm:"type:varchar;not null"`
	Email     string `json:"email" gorm:"type:varchar;unique;not null"`
	Bio       string `json:"bio" gorm:"type:varchar; not null"`
	Image     string `json:"image" gorm:"type:varchar; not null"`
	Password  string `json:"password" gorm:"type:varchar;not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelUser) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Password = utils.HashPassword(m.Password)
	m.CreatedAt = time.Now()

	return nil
}

func (m *ModelUser) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Password = utils.HashPassword(m.Password)
	m.UpdatedAt = time.Now()
	return nil
}
