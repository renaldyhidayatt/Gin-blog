package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityCategory interface {
	EntityCreate(input *schemas.SchemaCategories) (*models.ModelCategory, error)
	EntityResult(input *schemas.SchemaCategories) (*models.ModelCategory, error)
	EntityResults() (*[]models.ModelCategory, error)
	EntityDelete(input *schemas.SchemaCategories) (*models.ModelCategory, error)
	EntityUpdate(input *schemas.SchemaCategories) (*models.ModelCategory, error)
}
