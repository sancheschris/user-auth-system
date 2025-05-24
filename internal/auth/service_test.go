package auth

import (
	"testing"

	// "github.com/golang/mock/gomock"
	"github.com/sancheschris/user-auth-system/internal/model"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)


func TestRegister_Success(t *testing.T) {
	// Mock Controller
	ctrl := gomock.NewController(t)

	// Mock Object
	mockUserRepo := NewMockUserRepository(ctrl)
	service := NewAuthService(mockUserRepo)

	user := model.User{Username: "john", Password: "password123"}

	mockUserRepo.
		EXPECT().
		FindByUsername("john").
		Return(nil, nil)

	mockUserRepo.
		EXPECT().
		SaveUser(gomock.Any()).
		Return(nil)

	err := service.Register(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRegister_UserExists(t *testing.T) {
	// Mock Controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock Object
	mockUserRepo := NewMockUserRepository(ctrl)
	service := NewAuthService(mockUserRepo)

	user := model.User{Username: "john", Password: "password123"}

	mockUserRepo.
		EXPECT().
		FindByUsername("john").
		Return(&user, nil)

	err := service.Register(user)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	assert.EqualError(t, err, "user already exists")
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockUserRepository(ctrl)
	service := NewAuthService(mockUserRepo)

	// hash password for test
	hashed, _ := hashPassword("password123")
	existingUser := &model.User{Username: "john", Password: hashed}
	user := model.User{Username: "john", Password: "password123"}

	mockUserRepo.
		EXPECT().
		FindByUsername("john").
		Return(existingUser, nil)

	token, err := service.Login(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLogin_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockUserRepository(ctrl)
	service := NewAuthService(mockUserRepo)

	user := model.User{Username: "john", Password: "password123"}

	mockUserRepo.
		EXPECT().
		FindByUsername("john").
		Return(nil, assert.AnError)

	token, err := service.Login(user)
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.EqualError(t, err, "user not found")
}

func TestLogin_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockUserRepository(ctrl)
	service := NewAuthService(mockUserRepo)

	hashed, _ := hashPassword("password123")
	existingUser := &model.User{Username: "john", Password: hashed}
	user := model.User{Username: "john", Password: "wrongpassword"}

	mockUserRepo.
		EXPECT().
		FindByUsername("john").
		Return(existingUser, nil)

	token, err := service.Login(user)
	assert.Error(t, err)
	assert.Equal(t, "", token)
	assert.EqualError(t, err, "invalid password")
}
