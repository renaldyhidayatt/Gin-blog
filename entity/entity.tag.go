package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityTag interface {
	EntityCreate(input *schemas.SchemaTag) (*models.ModelTag, error)
	EntityResult(input *schemas.SchemaTag) (*models.ModelTag, error)
	EntityResults() (*[]models.ModelTag, error)
	EntityDelete(input *schemas.SchemaTag) (*models.ModelTag, error)
	EntityUpdate(input *schemas.SchemaTag) (*models.ModelTag, error)
}
