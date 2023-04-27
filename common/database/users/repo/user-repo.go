package user_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
)

func CreateUserByID(userID uint) error {
	tx := database.Users.Create(&user_entity.User{ID: userID})
	return tx.Error
}

func DeleteUserByID(userID uint) error {
	tx := database.Users.Delete(&user_entity.User{ID: userID})
	return tx.Error
}
