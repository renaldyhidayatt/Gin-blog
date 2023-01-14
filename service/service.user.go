package service

import (
	"ginBlog/entity"
	"ginBlog/models"
	"ginBlog/schemas"
)

type serviceUser struct {
	user entity.EntityUser
}

func NewServiceUser(user entity.EntityUser) *serviceUser {
	return &serviceUser{user: user}
}

func (s *serviceUser) EntityResults() (*[]models.ModelUser, error) {
	res, err := s.user.EntityResults()

	return res, err
}

func (s *serviceUser) EntityCreate(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user schemas.SchemasUser

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Email = input.Email
	user.Password = input.Password

	res, err := s.user.EntityCreate(&user)

	return res, err
}

func (s *serviceUser) EntityResult(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user schemas.SchemasUser
	user.ID = input.ID

	res, err := s.user.EntityResult(&user)

	return res, err

}

func (s *serviceUser) EntityUpdate(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user schemas.SchemasUser

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Image = input.Image
	user.Email = input.Email
	user.Password = input.Password

	res, err := s.user.EntityUpdate(&user)

	return res, err
}

func (s *serviceUser) EntityDelete(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user schemas.SchemasUser
	user.ID = input.ID

	res, err := s.user.EntityDelete(&user)

	return res, err

}
