package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db: db}
}

func (r *repositoryCategory) EntityCreate(input *schemas.SchemaCategories) (*models.ModelCategory, error) {
	var category models.ModelCategory
	category.Name = input.Name
	category.Description = input.Description

	db := r.db.Model(&category)

	addCategory := db.Debug().Create(&category).Commit()

	if addCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed name category already exit")
	}

	return &category, nil
}

func (r *repositoryCategory) EntityResults() (*[]models.ModelCategory, error) {
	var category []models.ModelCategory

	db := r.db.Model(&category)

	checkCategory := db.Debug().Find(&category)

	if checkCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed category not found")
	}

	return &category, nil
}

func (r *repositoryCategory) EntityResult(input *schemas.SchemaCategories) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.ID = input.ID

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category)

	if checkCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed result category")
	}

	return &category, nil
}

func (r *repositoryCategory) EntityUpdate(input *schemas.SchemaCategories) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.ID = input.ID

	db := r.db.Model(&category)

	checkCategoryID := db.Debug().First(&category)

	if checkCategoryID.RowsAffected < 1 {

		return &category, fmt.Errorf("failed result category")
	}

	category.Name = input.Name
	category.Description = input.Description

	updateCategory := db.Debug().Updates(&category)

	if updateCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed update category")
	}

	return &category, nil
}

func (r *repositoryCategory) EntityDelete(input *schemas.SchemaCategories) (*models.ModelCategory, error) {
	var category models.ModelCategory
	category.ID = input.ID

	db := r.db.Model(&category)

	checkCategory := db.Debug().First(&category)

	if checkCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed result category")
	}

	deleteCategory := db.Debug().Delete(&category)

	if deleteCategory.RowsAffected < 1 {
		return nil, fmt.Errorf("failed delete category")
	}

	return &category, nil
}
