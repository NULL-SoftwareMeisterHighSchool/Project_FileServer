package user_repo

import (
	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database"
	user_entity "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/entity"
)

func CreateUserByID(userID uint) {
	database.Users.Create(&user_entity.User{ID: userID})
}

func DeleteUserByID(userID uint) {
	database.Users.Delete(&user_entity.User{ID: userID})
}
