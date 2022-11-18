package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityTag interface {
	EntityCreate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelTag, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError)
}
