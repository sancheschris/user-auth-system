package auth

import (
	"fmt"

	"github.com/sancheschris/user-auth-system/internal/model"
	"github.com/sancheschris/user-auth-system/pkg/jwt"
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

func (s *AuthService) Login(user model.User) (string, error) {
	existingUser, err := s.Repo.FindByUsername(user.Username)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}
	if !checkPassword(user.Password, existingUser.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := jwt.GenerateToken(existingUser.Username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

