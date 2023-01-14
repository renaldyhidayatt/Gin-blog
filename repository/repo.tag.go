package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"

	"gorm.io/gorm"
)

type repositoryTag struct {
	db *gorm.DB
}

func NewRepositoryTag(db *gorm.DB) *repositoryTag {
	return &repositoryTag{db: db}
}

func (r *repositoryTag) EntityCreate(input *schemas.SchemaTag) (*models.ModelTag, error) {
	var tag models.ModelTag

	tag.Name = input.Name
	tag.Description = input.Description

	db := r.db.Model(&tag)

	checkName := db.Debug().First(&tag, "name = ?", input.Name)

	if checkName.RowsAffected > 0 {
		return nil, fmt.Errorf("failed tag: already name")
	}

	addTag := db.Debug().Create(&tag).Commit()

	if addTag.RowsAffected < 1 {
		return nil, fmt.Errorf("failed create tag")
	}

	return &tag, nil
}

func (r *repositoryTag) EntityResults() (*[]models.ModelTag, error) {
	var tag []models.ModelTag

	db := r.db.Model(&tag)

	checkTag := db.Debug().Find(&tag)

	if checkTag.RowsAffected < 1 {
		return nil, fmt.Errorf("failed results tag")
	}
	return &tag, nil
}

func (r *repositoryTag) EntityResult(input *schemas.SchemaTag) (*models.ModelTag, error) {
	var tag models.ModelTag

	tag.ID = input.ID

	db := r.db.Model(&tag)

	checkTag := db.Debug().First(&tag)

	if checkTag.RowsAffected < 1 {
		return nil, fmt.Errorf("failed result tag not found")
	}

	return &tag, nil
}

func (r *repositoryTag) EntityUpdate(input *schemas.SchemaTag) (*models.ModelTag, error) {
	var tag models.ModelTag

	tag.ID = input.ID

	db := r.db.Model(&tag)

	checTagID := db.Debug().First(&tag)

	if checTagID.RowsAffected < 1 {

		return nil, fmt.Errorf("failed query result tag not found")
	}

	tag.Name = input.Name
	tag.Description = input.Description

	updateTag := db.Debug().Updates(&tag)

	if updateTag.RowsAffected < 1 {

		return nil, fmt.Errorf("failed query update tag")
	}

	return &tag, nil
}

func (r *repositoryTag) EntityDelete(input *schemas.SchemaTag) (*models.ModelTag, error) {
	var tag models.ModelTag
	tag.ID = input.ID

	db := r.db.Model(&tag)

	checkTag := db.Debug().First(&tag)

	if checkTag.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query result: tag not found")
	}

	deleteTag := db.Debug().Delete(&tag)

	if deleteTag.RowsAffected < 1 {
		return nil, fmt.Errorf("failed query delete tag")
	}

	return &tag, nil
}
