package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityUser interface {
	EntityCreate(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelUser, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError)
}
