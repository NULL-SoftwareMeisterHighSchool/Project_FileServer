package database

import (
	"fmt"
	"log"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
	image_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/images/entity"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_NAME,
	)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}

	db = d

	migrateDB(db)
}

func migrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(
		&user_entity.User{},
		&article_entity.Article{},
		&article_entity.ArticleBody{},
		&image_entity.Image{},
		&comment_entity.Comment{},
	); err != nil {
		log.Fatalf("cannot migrate db: %v", err)
	}
}

func Articles() *gorm.DB {
	return db.Model(&article_entity.Article{})
}

func Images() *gorm.DB {
	return db.Model(&image_entity.Image{})
}

func ArticleBodies() *gorm.DB {
	return db.Model(&article_entity.ArticleBody{})
}

func Users() *gorm.DB {
	return db.Model(&user_entity.User{})
}

func Comments() *gorm.DB {
	return db.Model(&comment_entity.Comment{})
}

func ArticleLikes() *gorm.DB {
	return db.Table("user_likes")
}

func DB() *gorm.DB {
	return db
}
