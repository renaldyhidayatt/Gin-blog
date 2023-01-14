package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityArticle interface {
	EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityResults() (*[]models.ModelArticle, error)
	EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, error)
}
