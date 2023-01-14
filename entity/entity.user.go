package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityUser interface {
	EntityCreate(input *schemas.SchemasUser) (*models.ModelUser, error)
	EntityResult(input *schemas.SchemasUser) (*models.ModelUser, error)
	EntityResults() (*[]models.ModelUser, error)
	EntityDelete(input *schemas.SchemasUser) (*models.ModelUser, error)
	EntityUpdate(input *schemas.SchemasUser) (*models.ModelUser, error)
}
