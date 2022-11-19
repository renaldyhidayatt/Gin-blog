package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceComment struct {
	comment entity.EntityComment
}

func NewServiceComment(comment entity.EntityComment) *serviceComment {
	return &serviceComment{comment: comment}
}

func (s *serviceComment) EntityResults() (*[]models.ModelComment, schemas.SchemaDatabaseError) {
	res, err := s.comment.EntityResults()

	return res, err
}

func (s *serviceComment) EntityCreate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment schemas.SchemaComment

	comment.Content = input.Content
	comment.UserID = input.UserID

	res, err := s.comment.EntityCreate(&comment)

	return res, err
}

func (s *serviceComment) EntityResult(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment schemas.SchemaComment

	comment.ID = input.ID

	res, err := s.comment.EntityResult(&comment)

	return res, err
}

func (s *serviceComment) EntityUpdate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment schemas.SchemaComment

	comment.Content = input.Content
	comment.UserID = input.UserID

	res, err := s.comment.EntityUpdate(&comment)

	return res, err
}

func (s *serviceComment) EntityDelete(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment schemas.SchemaComment

	comment.ID = input.ID

	res, err := s.comment.EntityDelete(&comment)

	return res, err
}
