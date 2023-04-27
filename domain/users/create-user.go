package users

import user_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/repo"

func CreateUser(userID uint) error {
	return user_repo.CreateUserByID(userID)
}
