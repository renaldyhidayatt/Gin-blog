package entity

import (
	"ginBlog/models"
	"ginBlog/schemas"
)

type EntityArticle interface {
	EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelArticle, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError)
}
