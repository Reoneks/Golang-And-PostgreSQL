package user

import (
	"errors"
	"strconv"
	"strings"
	"test/utils"
)

type UserService interface {
	GetUser(userId int64) (*User, error)
	GetUserByEmail(username string) (*User, error)
	Registration(user UserDto) (*User, error)
	GetUsers(filter UserFilter) ([]User, error)
	DeleteUser(user int64) error
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func (s *UserServiceImpl) GetUser(userId int64) (*User, error) {
	result, err := s.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDto(*result)
	return &resultUser, nil
}

func (s *UserServiceImpl) Registration(user UserDto) (*User, error) {
	password, err := utils.Encrypt(user.Password)
	user.Password = password
	if err != nil {
		return nil, err
	}
	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDto(*result)
	return &resultUser, nil
}
func (s *UserServiceImpl) GetUsers(filter UserFilter) ([]User, error) {
	var search []string
	if filter.Email != "" {
		search = append(search, "email LIKE '%"+filter.Email+"%'")
	}
	if filter.FirstName != "" {
		search = append(search, "first_name LIKE '%"+filter.FirstName+"%'")
	}
	if filter.LastName != "" {
		search = append(search, "last_name LIKE '%"+filter.LastName+"%'")
	}
	if filter.Status != 0 {
		search = append(search, "status = '"+strconv.FormatInt(filter.Status, 10)+"'")
	}
	result, err := s.userRepository.GetUsers(strings.Join(search, " AND "))
	if err != nil {
		return nil, err
	}
	resultUser := FromUserDtos(result)
	return resultUser, nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*User, error) {
	someUsers, err := s.userRepository.GetUsers("email = '" + email + "'")
	if err != nil {
		return nil, err
	} else if len(someUsers) == 0 {
		return nil, errors.New("user not found")
	}
	resultUser := FromUserDto(someUsers[0])
	return &resultUser, nil
}

func (s *UserServiceImpl) DeleteUser(user int64) error {
	return s.userRepository.DeleteUser(user)
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{
		userRepository,
	}
}
