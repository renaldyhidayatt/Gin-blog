package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceTag struct {
	tag entity.EntityTag
}

func NewServiceTag(tag entity.EntityTag) *serviceTag {
	return &serviceTag{tag: tag}
}

func (s *serviceTag) EntityResults() (*[]models.ModelTag, schemas.SchemaDatabaseError) {
	res, err := s.tag.EntityResults()
	return res, err
}

func (s *serviceTag) EntityCreate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag schemas.SchemaTag

	tag.Name = input.Name
	tag.Description = input.Description

	res, err := s.tag.EntityCreate(&tag)

	return res, err
}

func (s *serviceTag) EntityResult(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag schemas.SchemaTag

	tag.ID = input.ID

	res, err := s.tag.EntityResult(&tag)

	return res, err
}

func (s *serviceTag) EntityUpdate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag schemas.SchemaTag

	tag.Name = input.Name
	tag.Description = input.Description

	res, err := s.tag.EntityUpdate(&tag)

	return res, err
}

func (s *serviceTag) EntityDelete(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag schemas.SchemaTag

	tag.ID = input.ID

	res, err := s.tag.EntityDelete(&tag)

	return res, err
}
