package database

import (
	"fmt"
	"log"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	article_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/articles/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var Articles = db.Model(article_entity.New())

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_NAME,
	)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}
	d.AutoMigrate(article_entity.New())
	db = d
}
