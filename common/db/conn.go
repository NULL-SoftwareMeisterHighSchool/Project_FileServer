package db

import (
	"fmt"
	"log"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
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
}
