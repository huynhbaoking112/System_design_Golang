package service

import (
	"github.com/huynhbaoking112/System_design_Golang/internal/repo"
	"github.com/huynhbaoking112/System_design_Golang/package/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserRepoService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// // user repo ur
// func (us *UserService) GetInfoUserService() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

func NewUserService(
	userRepo repo.IUserRepo,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {

	//1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSusscess
}
