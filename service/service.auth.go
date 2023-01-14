package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceAuth struct {
	auth entity.EntityAuth
}

func NewServiceAuth(auth entity.EntityAuth) *serviceAuth {
	return &serviceAuth{auth: auth}
}

func (s *serviceAuth) EntityRegister(input *schemas.SchemaAuth) (*models.ModelUser, error) {
	var schema schemas.SchemaAuth
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Bio = input.Bio
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityRegister(&schema)
	return res, err
}

func (s *serviceAuth) EntityLogin(input *schemas.SchemaAuth) (*models.ModelUser, error) {
	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.auth.EntityLogin(&schema)

	return res, err
}
