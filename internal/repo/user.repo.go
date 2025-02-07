package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // user  repo u
// func (ur *UserRepo) GetInfoUser() string {
// 	return "Tipjs"
// }

//INTERFACE VERSION

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct {
}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepo {
	return &userRepo{}
}
