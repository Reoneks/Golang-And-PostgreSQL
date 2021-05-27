package test

import (
	"errors"
	"test/auth"
	. "test/config"
	"test/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

var authServiceSuccess auth.AuthService
var authServiceError auth.AuthService

type UserServiceImplSuccess struct {
	userRepository user.UserRepository
}

func (s *UserServiceImplSuccess) GetUser(userId int64) (*user.User, error) {
	return nil, nil
}
func (s *UserServiceImplSuccess) Registration(nUser user.UserDto) (*user.User, error) {
	return nil, nil
}
func (s *UserServiceImplSuccess) GetUsers(search string) ([]user.User, error) {
	return nil, nil
}
func (s *UserServiceImplSuccess) GetUserByEmail(email string) (*user.User, error) {
	return &user.User{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "$2a$10$wWobk.2ScNN2nztxToTBnuEh7TwD1qpWwwPnY8baeLxDeSSMulEOq",
	}, nil
}
func (s *UserServiceImplSuccess) DeleteUser(user int64) error {
	return s.userRepository.DeleteUser(user)
}

type UserServiceImplError struct {
	userRepository user.UserRepository
}

func (s *UserServiceImplError) GetUser(userId int64) (*user.User, error) {
	return nil, nil
}
func (s *UserServiceImplError) Registration(nUser user.UserDto) (*user.User, error) {
	return nil, nil
}
func (s *UserServiceImplError) GetUsers(search string) ([]user.User, error) {
	return nil, nil
}
func (s *UserServiceImplError) GetUserByEmail(email string) (*user.User, error) {
	return nil, errors.New("Some error")
}
func (s *UserServiceImplError) DeleteUser(user int64) error {
	return s.userRepository.DeleteUser(user)
}

func init() {
	config := NewConfig()
	authServiceSuccess = auth.NewAuthService(&UserServiceImplSuccess{}, config.JWT())
	authServiceError = auth.NewAuthService(&UserServiceImplError{}, config.JWT())
}

func TestLoginSuccess(t *testing.T) {
	result, err := authServiceSuccess.Login("Kendall22@gmail.com", "it4xxOEvrY2vua7")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(
		t,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.jYyRJbb0WImFoUUdcslQQfwnXTHJzne-6tsPd8Hrw0I",
		result.Token,
	)
}
func TestLoginError(t *testing.T) {
	result, err := authServiceError.Login("Kendall22@gmail.com", "it4xxOEvrY2vua7")
	if err == nil && result != nil {
		t.Error("Error can`t be nil here or result != nil")
	}
	assert.Equal(t, err.Error(), "Some error")
}
