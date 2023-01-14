package repository

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/schemas"

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

func (r *repositoryArticle) EntityCreate(input *schemas.SchemaArticle) (*models.ModelArticle, error) {
	var article models.ModelArticle

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article, "title = ?", input.Title)

	if checkArticle.RowsAffected > 0 {
		return nil, fmt.Errorf("already name title article")
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

		return nil, fmt.Errorf("failed create article")
	}

	return &article, nil

}

func (r *repositoryArticle) EntityResults(page, size int) (*[]models.ModelArticle, error) {
	var article []models.ModelArticle

	db := r.db.Model(&article)

	checkArticle := db.Debug().Limit(size).Offset((page - 1) * size).Find(&article)

	if checkArticle.RowsAffected < 1 {
		return nil, fmt.Errorf("not found articles")
	}
	return &article, nil
}

func (r *repositoryArticle) EntityCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.ModelArticle{}).Count(&count).Error
	return count, err
}

func (r *repositoryArticle) EntityResult(input *schemas.SchemaArticle) (*models.ModelArticle, error) {
	var article models.ModelArticle

	article.ID = input.ID

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article)

	if checkArticle.RowsAffected < 1 {
		return nil, fmt.Errorf("article not found")
	}

	return &article, nil
}

func (r *repositoryArticle) EntityUpdate(input *schemas.SchemaArticle) (*models.ModelArticle, error) {
	var article models.ModelArticle

	article.ID = input.ID

	db := r.db.Model(&article)

	checkArticleID := db.Debug().First(&article)

	if checkArticleID.RowsAffected < 1 {

		return nil, fmt.Errorf("article not found")
	}

	article.Title = input.Title
	article.Description = input.Description
	article.Body = input.Body
	article.UserID = input.UserID

	updateArticle := db.Debug().Updates(&article)

	if updateArticle.RowsAffected < 1 {

		return nil, fmt.Errorf("failed update article")
	}

	return &article, nil
}

func (r *repositoryArticle) EntityDelete(input *schemas.SchemaArticle) (*models.ModelArticle, error) {
	var article models.ModelArticle

	article.ID = input.ID

	db := r.db.Model(&article)

	checkArticle := db.Debug().First(&article)

	if checkArticle.RowsAffected < 1 {

		return nil, fmt.Errorf("article not found")
	}

	deleteArticle := db.Debug().Delete(&article)

	if deleteArticle.RowsAffected < 1 {

		return nil, fmt.Errorf("failed delete article")
	}

	return &article, nil

}
