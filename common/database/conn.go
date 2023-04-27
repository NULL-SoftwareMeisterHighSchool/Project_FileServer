package database

import (
	"fmt"
	"log"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	comment_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/comments/entity"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var Articles = db.Model(&article_entity.Article{})
var Users = db.Model(&user_entity.User{})
var Comments = db.Model(&comment_entity.Comment{})

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
	database.AutoMigrate(&article_entity.Article{})
	database.AutoMigrate(&article_entity.ArticleBody{})
	database.AutoMigrate(&user_entity.User{})
	database.AutoMigrate(&comment_entity.Comment{})
}
