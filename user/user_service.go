package user

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type UserService interface {
	GetUser(userId int64) (*UserDto, error)
	Registration(user UserDto) (*UserDto, error)
	GetUsers(search string) ([]*UserDto, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func (s *UserServiceImpl) GetUser(userId int64) (*UserDto, error) {
	return s.userRepository.GetUser(userId)
}

func encrypt(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha, nil
}

func (s *UserServiceImpl) Registration(user UserDto) (*UserDto, error) {
	password, err := encrypt(user.Password)
	user.Password = password
	if err != nil {
		return nil, err
	}
	return s.userRepository.CreateUser(user)
}
func (s *UserServiceImpl) GetUsers(search string) ([]*UserDto, error) {
	params := strings.Split(search, ":")
	if len(params) == 2 {
		return s.userRepository.GetUsers(params[0] + " LIKE '%" + params[1] + "%'")
	}
	return s.userRepository.GetUsers("")
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{
		userRepository,
	}
}
