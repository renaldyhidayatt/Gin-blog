package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryTag struct {
	db *gorm.DB
}

func NewRepositoryTag(db *gorm.DB) *repositoryTag {
	return &repositoryTag{db: db}
}

func (r *repositoryTag) EntityCreate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag models.ModelTag

	tag.Name = input.Name
	tag.Description = input.Description

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&tag)

	checkName := db.Debug().First(&tag, "name = ?", input.Name)

	if checkName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
	}

	addTag := db.Debug().Create(&tag).Commit()

	if addTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &tag, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &tag, <-err
}

func (r *repositoryTag) EntityResults() (*[]models.ModelTag, schemas.SchemaDatabaseError) {
	var tag []models.ModelTag

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&tag)

	checkTag := db.Debug().Find(&tag)

	if checkTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}
	return &tag, <-err
}

func (r *repositoryTag) EntityResult(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag models.ModelTag

	tag.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&tag)

	checkTag := db.Debug().First(&tag)

	if checkTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &tag, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &tag, <-err
}

func (r *repositoryTag) EntityUpdate(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag models.ModelTag

	tag.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&tag)

	checTagID := db.Debug().First(&tag)

	if checTagID.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &tag, <-err
	}

	tag.Name = input.Name
	tag.Description = input.Description

	updateTag := db.Debug().Updates(&tag)

	if updateTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &tag, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &tag, <-err
}

func (r *repositoryTag) EntityDelete(input *schemas.SchemaTag) (*models.ModelTag, schemas.SchemaDatabaseError) {
	var tag models.ModelTag
	tag.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&tag)

	checkTag := db.Debug().First(&tag)

	if checkTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &tag, <-err
	}

	deleteTag := db.Debug().Delete(&tag)

	if deleteTag.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &tag, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &tag, <-err
}
