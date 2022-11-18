package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) EntityCreate(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &user, <-err
}

func (r *repositoryUser) EntityResults() (*[]models.ModelUser, schemas.SchemaDatabaseError) {
	var user []models.ModelUser

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUser := db.Debug().Find(&user)

	if checkUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}
	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *repositoryUser) EntityResult(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser

	user.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *repositoryUser) EntityUpdate(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser

	user.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUserId := db.Debug().First(&user)

	if checkUserId.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &user, <-err
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Image = input.Image
	user.Email = input.Email
	user.Password = input.Password

	updateUser := db.Debug().Updates(&user)

	if updateUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &user, <-err

}

func (r *repositoryUser) EntityDelete(input *schemas.SchemasUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &user, <-err
	}

	deleteUser := db.Debug().Delete(&user)

	if deleteUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}

		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &user, <-err
}
