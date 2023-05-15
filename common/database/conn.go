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

var (
	Articles      *gorm.DB
	Images        *gorm.DB
	ArticleBodies *gorm.DB
	Users         *gorm.DB
	Comments      *gorm.DB
	ArticleLikes  *gorm.DB
	DB            *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}

	DB = db

	migrateDB(db)
	loadTables(db)
}

func loadTables(db *gorm.DB) {
	Articles = db.Model(&article_entity.Article{})
	Images = db.Model(&image_entity.Image{})
	ArticleBodies = db.Model(&article_entity.ArticleBody{})
	Users = db.Model(&user_entity.User{})
	Comments = db.Model(&comment_entity.Comment{})
	ArticleLikes = db.Table("user_likes")
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
