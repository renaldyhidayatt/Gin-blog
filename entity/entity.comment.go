package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityComment interface {
	EntityCreate(input *schemas.SchemaComment) (*models.ModelComment, error)
	EntityResult(input *schemas.SchemaComment) (*models.ModelComment, error)
	EntityResults() (*[]models.ModelComment, error)
	EntityDelete(input *schemas.SchemaComment) (*models.ModelComment, error)
	EntityUpdate(input *schemas.SchemaComment) (*models.ModelComment, error)
}
