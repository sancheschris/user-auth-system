package auth

import (
	"fmt"

	"github.com/sancheschris/user-auth-system/internal/model"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct {
	Repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(user model.User) error {
	existingUser, _ := s.Repo.FindByUsername(user.Username)
	if existingUser != nil {
		return fmt.Errorf("user already exists")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	newUser := &model.User{
		Username: user.Username,
		Password: hashedPassword,
	}

	return s.Repo.SaveUser(newUser)
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}