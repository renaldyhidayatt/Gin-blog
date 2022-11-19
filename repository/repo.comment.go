package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryComment struct {
	db *gorm.DB
}

func NewRepositoryComment(db *gorm.DB) *repositoryComment {
	return &repositoryComment{db: db}
}

func (r *repositoryComment) EntityCreate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment models.ModelComment
	comment.Content = input.Content
	comment.UserId = input.UserID
	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&comment)

	addComment := db.Debug().Create(&comment).Commit()

	if addComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_01",
		}

		return &comment, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &comment, <-err
}

func (r *repositoryComment) EntityResults() (*[]models.ModelComment, schemas.SchemaDatabaseError) {
	var comment []models.ModelComment

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&comment)

	checkComment := db.Debug().Find(&comment)

	if checkComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}
	return &comment, <-err
}

func (r *repositoryComment) EntityResult(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment models.ModelComment

	comment.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&comment)

	checkComment := db.Debug().First(&comment)

	if checkComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &comment, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &comment, <-err
}

func (r *repositoryComment) EntityUpdate(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment models.ModelComment

	comment.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&comment)

	checkCommentID := db.Debug().First(&comment)

	if checkCommentID.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &comment, <-err
	}

	comment.Content = input.Content

	updateComment := db.Debug().Updates(&comment)

	if updateComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &comment, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &comment, <-err
}

func (r *repositoryComment) EntityDelete(input *schemas.SchemaComment) (*models.ModelComment, schemas.SchemaDatabaseError) {
	var comment models.ModelComment
	comment.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&comment)

	checkComment := db.Debug().First(&comment)

	if checkComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &comment, <-err
	}

	deleteComment := db.Debug().Delete(&comment)

	if deleteComment.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &comment, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &comment, <-err
}
