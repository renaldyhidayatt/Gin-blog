package config

import (
	"fmt"
	"ginBlog/models"
	"ginBlog/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.GodotEnv("DB_HOST"),
		utils.GodotEnv("DB_PORT"),
		utils.GodotEnv("DB_USER"),
		utils.GodotEnv("DB_PASS"),
		utils.GodotEnv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err)
		return nil
	}

	err = db.AutoMigrate(&models.ModelUser{}, &models.ModelArticle{}, &models.ModelCategory{}, &models.ModelComment{}, &models.ModelTag{})
	if err != nil {
		defer logrus.Info("Database Connection failed")
		logrus.Fatal(err)
		return nil
	}

	return db

}
