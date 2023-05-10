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
var Articles = db.Model(&article_entity.Article{})
var Images = db.Model(&image_entity.Image{})
var ArticleBodies = db.Model(&article_entity.ArticleBody{})
var Users = db.Model(&user_entity.User{})
var Comments = db.Model(&comment_entity.Comment{})
var ArticleLikes = db.Table("user_likes")

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

func migrateDB(database *gorm.DB) {
	Articles.AutoMigrate()
	ArticleBodies.AutoMigrate()
	Users.AutoMigrate()
	Comments.AutoMigrate()
	Images.AutoMigrate()
}
