package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityArticle interface {
	EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityResults(page, size int) (*[]models.ModelArticle, error)
	EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, error)
	EntityCount() (int64, error)
}
