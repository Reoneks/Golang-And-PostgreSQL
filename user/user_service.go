package user

import (
	"errors"
	"strings"
	"test/utils"
)

type UserService interface {
	GetUser(userId int64) (*UserDto, error)
	GetUserByUsername(username string) (*UserDto, error)
	Registration(user UserDto) (*UserDto, error)
	GetUsers(search string) ([]*UserDto, error)
	DeleteUser(user int64) error
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func (s *UserServiceImpl) GetUser(userId int64) (*UserDto, error) {
	return s.userRepository.GetUser(userId)
}

func (s *UserServiceImpl) Registration(user UserDto) (*UserDto, error) {
	password, err := utils.Encrypt(user.Password)
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

func (s *UserServiceImpl) GetUserByUsername(username string) (*UserDto, error) {
	params := strings.Split(username, " ")
	if len(params) != 2 {
		return nil, errors.New("incorrect username")
	}
	someUsers, err := s.userRepository.GetUsers("first_name = '" + params[0] + "' AND last_name = '" + params[1] + "'")
	if err != nil {
		return nil, err
	}
	return someUsers[0], nil
}

func (s *UserServiceImpl) DeleteUser(user int64) error {
	return s.userRepository.DeleteUser(user)
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{
		userRepository,
	}
}
