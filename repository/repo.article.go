package repository

import (
	"ginBlog/models"
	"ginBlog/schemas"
	"net/http"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type repositoryArticle struct {
	db *gorm.DB
}

func NewRepositoryArticle(db *gorm.DB) *repositoryArticle {
	return &repositoryArticle{
		db: db,
	}
}

func (r *repositoryArticle) EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article models.ModelArticle

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article, "title = ?", input.Title)

	if checkArticle.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
	}

	tags := make([]models.ModelTag, len(input.Tag))
	categories := make([]models.ModelCategory, len(input.Categories))

	for i, tag := range input.Tag {
		r.db.Where(&models.ModelTag{Slug: slug.Make(tag.Name)}).Attrs(models.ModelTag{Name: tag.Name, Description: tag.Description}).FirstOrCreate(&tags[i])
	}

	for i := range input.Categories {
		r.db.Where(&models.ModelCategory{Slug: slug.Make(input.Categories[i].Name)}).Attrs(models.ModelCategory{Name: input.Categories[i].Name, Description: input.Categories[i].Description}).FirstOrCreate(&categories[i])
	}

	article.Tags = tags
	article.Categories = categories

	addArticle := db.Debug().Create(&article).Commit()

	if addArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &article, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &article, <-err

}

func (r *repositoryArticle) EntityResults() (*[]models.ModelArticle, schemas.SchemaDatabaseError) {
	var article []models.ModelArticle

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&article)

	checkArticle := db.Debug().Find(&article)

	if checkArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}
	err <- schemas.SchemaDatabaseError{}
	return &article, <-err
}

func (r *repositoryArticle) EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article models.ModelArticle

	article.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article)

	if checkArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &article, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &article, <-err
}

func (r *repositoryArticle) EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article models.ModelArticle

	article.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&article)

	checkArticleID := db.Debug().First(&article)

	if checkArticleID.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &article, <-err
	}

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID

	updateArticle := db.Debug().Updates(&article)

	if updateArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &article, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &article, <-err
}

func (r *repositoryArticle) EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, schemas.SchemaDatabaseError) {
	var article models.ModelArticle

	article.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article)

	if checkArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &article, <-err
	}

	deleteArticle := db.Debug().Delete(&article)

	if deleteArticle.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &article, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &article, <-err

}
