package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityAuth interface {
	EntityRegister(input *schemas.SchemaAuth) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemaAuth) (*models.ModelUser, schemas.SchemaDatabaseError)
}
