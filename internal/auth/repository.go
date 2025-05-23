package auth

import "github.com/sancheschris/user-auth-system/internal/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	SaveUser(user *model.User) error
}