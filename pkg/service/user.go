package service

import "github.com/chrismalcolm/recipe/pkg/model"

type UserAdmin interface {
	AddUser(user model.User) (userID model.UserID, err error)
	EditUser(user model.User) (err error)
	DeleteUser(userID model.UserID) (err error)
}

type UserService struct {
	ua UserAdmin
}

func NewUserService(userAdmin UserAdmin) (us *UserService) {
	return &UserService{
		ua: userAdmin,
	}
}
