package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceCategory struct {
	category entity.EntityCategory
}

func NewServiceCategory(category entity.EntityCategory) *serviceCategory {
	return &serviceCategory{category: category}
}

func (s *serviceCategory) EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError) {
	res, err := s.category.EntityResults()

	return res, err
}

func (s *serviceCategory) EntityCreate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategories

	category.Name = input.Name
	category.Description = input.Description

	res, err := s.category.EntityCreate(&category)

	return res, err
}

func (s *serviceCategory) EntityResult(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategories

	category.ID = input.ID

	res, err := s.category.EntityResult(&category)

	return res, err
}

func (s *serviceCategory) EntityUpdate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategories

	category.Name = input.Name
	category.Description = input.Description

	res, err := s.category.EntityUpdate(&category)

	return res, err
}

func (s *serviceCategory) EntityDelete(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategories

	category.ID = input.ID

	res, err := s.category.EntityDelete(&category)

	return res, err
}
