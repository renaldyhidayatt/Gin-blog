package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"

	"gorm.io/gorm"
)

type repositoryComment struct {
	db *gorm.DB
}

func NewRepositoryComment(db *gorm.DB) *repositoryComment {
	return &repositoryComment{db: db}
}

func (r *repositoryComment) EntityCreate(input *schemas.SchemaComment) (*models.ModelComment, error) {
	var comment models.ModelComment
	comment.Content = input.Content
	comment.UserId = input.UserID

	db := r.db.Model(&comment)

	_ = db.Debug().Create(&comment).Commit()

	return &comment, nil
}

func (r *repositoryComment) EntityResults() (*[]models.ModelComment, error) {
	var comment []models.ModelComment

	db := r.db.Model(&comment)

	checkComment := db.Debug().Find(&comment)

	if checkComment.RowsAffected < 1 {
		return nil, fmt.Errorf("failed results comments")
	}
	return &comment, nil
}

func (r *repositoryComment) EntityResult(input *schemas.SchemaComment) (*models.ModelComment, error) {
	var comment models.ModelComment

	comment.ID = input.ID

	db := r.db.Model(&comment)

	checkComment := db.Debug().First(&comment)

	if checkComment.RowsAffected < 1 {
		return nil, fmt.Errorf("failed result comment")
	}

	return &comment, nil
}

func (r *repositoryComment) EntityUpdate(input *schemas.SchemaComment) (*models.ModelComment, error) {
	var comment models.ModelComment

	comment.ID = input.ID

	db := r.db.Model(&comment)

	checkCommentID := db.Debug().First(&comment)

	if checkCommentID.RowsAffected < 1 {
		return nil, fmt.Errorf("failed result comment")
	}

	comment.Content = input.Content

	updateComment := db.Debug().Updates(&comment)

	if updateComment.RowsAffected < 1 {
		return nil, fmt.Errorf("failed update comment")
	}

	return &comment, nil
}

func (r *repositoryComment) EntityDelete(input *schemas.SchemaComment) (*models.ModelComment, error) {
	var comment models.ModelComment
	comment.ID = input.ID

	db := r.db.Model(&comment)

	checkComment := db.Debug().First(&comment)

	if checkComment.RowsAffected < 1 {

		return nil, fmt.Errorf("failed result comment")
	}

	deleteComment := db.Debug().Delete(&comment)

	if deleteComment.RowsAffected < 1 {
		return nil, fmt.Errorf("failed delete category not found")
	}

	return &comment, nil
}
