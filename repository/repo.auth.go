package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"
	"ginBlog/utils"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) EntityRegister(input *schemas.SchemaAuth) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		return nil, fmt.Errorf("failed query email already")
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {

		return nil, fmt.Errorf("failed create user")
	}

	return &user, nil
}

func (r *repositoryAuth) EntityLogin(input *schemas.SchemaAuth) (*models.ModelUser, error) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected < 1 {
		return nil, fmt.Errorf("email not found")
	}

	checkPasswordMatch := utils.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		return nil, fmt.Errorf("password salah")
	}
	return &user, nil
}
