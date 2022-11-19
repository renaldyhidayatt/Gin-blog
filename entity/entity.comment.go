package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityComment interface {
	EntityCreate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelComment, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError)
}
