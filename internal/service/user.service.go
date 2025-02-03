package service

import "github.com/huynhbaoking112/System_design_Golang/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserRepoService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user  repo ur
func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUser()
}
