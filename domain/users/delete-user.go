package users

import user_repo "github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/database/users/repo"

func DeleteUser(userID uint) error {
	return user_repo.DeleteUserByID(userID)
}
