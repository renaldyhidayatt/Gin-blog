package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityCategory interface {
	EntityCreate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError)
}
