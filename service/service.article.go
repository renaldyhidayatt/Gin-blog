package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceArticle struct {
	article entity.EntityArticle
}

func NewServideArticle(article entity.EntityArticle) *serviceArticle {
	return &serviceArticle{article: article}
}

func (s *serviceArticle) EntityResults() (*[]models.ModelArticle, schemas.SchemaDatabaseError) {
	res, err := s.article.EntityResults()

	return res, err
}

func (s *serviceArticle) EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article schemas.SchemaArticle

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID
	article.Tag = input.Tag
	article.Categories = input.Categories

	res, err := s.article.EntityCreate(&article)

	return res, err

}

func (s *serviceArticle) EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article schemas.SchemaArticle

	article.ID = input.ID

	res, err := s.article.EntityResult(&article)

	return res, err
}

func (s *serviceArticle) EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article schemas.SchemaArticle

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID

	res, err := s.article.EntityUpdate(&article)

	return res, err
}

func (s *serviceArticle) EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article schemas.SchemaArticle

	article.ID = input.ID

	res, err := s.article.EntityDelete(&article)

	return res, err
}
