package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) EntityCreate(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		return nil, fmt.Errorf("failed query user: already email")
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {

		return &user, fmt.Errorf("failed query user: create user")
	}

	return &user, nil
}

func (r *repositoryUser) EntityResults() (*[]models.ModelUser, error) {
	var user []models.ModelUser

	db := r.db.Model(&user)

	checkUser := db.Debug().Find(&user)

	if checkUser.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query user: result not found")
	}

	return &user, nil
}

func (r *repositoryUser) EntityResult(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {

		return &user, fmt.Errorf("failed query user: result not found")
	}

	return &user, nil
}

func (r *repositoryUser) EntityUpdate(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUserId := db.Debug().First(&user)

	if checkUserId.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query user: not found user")
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Bio = input.Bio
	user.Image = input.Image
	user.Email = input.Email
	user.Password = input.Password

	updateUser := db.Debug().Updates(&user)

	if updateUser.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query user: update user failed")
	}

	return &user, nil

}

func (r *repositoryUser) EntityDelete(input *schemas.SchemasUser) (*models.ModelUser, error) {
	var user models.ModelUser
	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query user: not found user")
	}

	deleteUser := db.Debug().Delete(&user)

	if deleteUser.RowsAffected < 1 {

		return nil, fmt.Errorf("failed query user: delete user failed")
	}

	return &user, nil
}
