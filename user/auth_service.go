package user

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Token string   `json:"token"`
	User  *UserDto `json:"user"`
}

type AuthService interface {
	Login(username, password string) (*Login, error)
}

type AuthServiceImpl struct {
	userService UserService
}

func compare(user UserDto, password string) error {
	userPassword, err := encrypt(password)
	if err != nil {
		return err
	}
	if user.Password != userPassword {
		return errors.New("entered the wrong password")
	}
	return nil
}

func (s *AuthServiceImpl) Login(username, password string) (*Login, error) {
	user, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	err = compare(*user, password)
	if err != nil {
		return nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
	})
	tokenString, err := token.SignedString([]byte("test-secret"))
	if err != nil {
		return nil, err
	}
	return &Login{
		Token: tokenString,
		User:  user,
	}, nil
}

func NewAuthService(userService UserService) AuthService {
	return &AuthServiceImpl{
		userService,
	}
}
