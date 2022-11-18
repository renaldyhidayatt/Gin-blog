package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"ginBlog/utils"
	"net/http"

	"gorm.io/gorm"
)

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *repositoryAuth {
	return &repositoryAuth{db: db}
}

func (r *repositoryAuth) EntityRegister(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Image = input.Image
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &user, <-err
}

func (r *repositoryAuth) EntityLogin(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_login_01",
		}
		return &user, <-err
	}

	checkPasswordMatch := utils.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}
