package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db: db}
}

func (r *repositoryCategory) EntityCreate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.Name = input.Name
	category.Description = input.Description

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkName := db.Debug().First(&category, "name = ?", input.Name)

	if checkName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
	}

	addCategory := db.Debug().Create(&category).Commit()

	if addCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &category, <-err
}

func (r *repositoryCategory) EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError) {
	var category []models.ModelCategory

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategory := db.Debug().Find(&category)

	if checkCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}
	return &category, <-err
}

func (r *repositoryCategory) EntityResult(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category)

	if checkCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &category, <-err
}

func (r *repositoryCategory) EntityUpdate(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryID := db.Debug().First(&category)

	if checkCategoryID.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &category, <-err
	}

	category.Name = input.Name
	category.Description = input.Description

	updateCategory := db.Debug().Updates(&category)

	if updateCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &category, <-err
}

func (r *repositoryCategory) EntityDelete(input *schemas.SchemaCategories) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory
	category.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category)

	if checkCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &category, <-err
	}

	deleteCategory := db.Debug().Delete(&category)

	if deleteCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &category, <-err
}
