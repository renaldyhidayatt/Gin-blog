package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
}
