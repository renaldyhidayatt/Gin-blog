package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemaAuth) (*models.ModelUser, error)
	EntityLogin(input *schemas.SchemaAuth) (*models.ModelUser, error)
}
