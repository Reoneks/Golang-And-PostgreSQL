package test

import (
	"errors"
	"test/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var userServiceSuccess user.UserService
var userServiceError user.UserService

type UserRepositoryImplSuccess struct{ db *gorm.DB }

func (r *UserRepositoryImplSuccess) GetUser(id int64) (*user.UserDto, error) {
	return &user.UserDto{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, nil
}
func (r *UserRepositoryImplSuccess) CreateUser(nUser user.UserDto) (*user.UserDto, error) {
	return &user.UserDto{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, nil
}
func (r *UserRepositoryImplSuccess) UpdateUser(nUser user.UserDto) (*user.UserDto, error) {
	return &user.UserDto{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, nil
}
func (r *UserRepositoryImplSuccess) DeleteUser(id int64) error {
	return nil
}
func (r *UserRepositoryImplSuccess) GetUsers(where string) ([]user.UserDto, error) {
	return []user.UserDto{
		{
			Id:        1,
			FirstName: "Thaddeus",
			LastName:  "Gerhold",
			Email:     "Kendall22@gmail.com",
			Password:  "it4xxOEvrY2vua7",
		},
		{
			Id:        2,
			FirstName: "Vernon",
			LastName:  "Hand",
			Email:     "Baylee0@hotmail.com",
			Password:  "1o0kLTPRTlYKgSv",
		},
	}, nil
}

type UserRepositoryImplError struct{ db *gorm.DB }

func (r *UserRepositoryImplError) GetUser(id int64) (*user.UserDto, error) {
	return nil, errors.New("Some error")
}
func (r *UserRepositoryImplError) CreateUser(user user.UserDto) (*user.UserDto, error) {
	return nil, errors.New("Some error")
}
func (r *UserRepositoryImplError) UpdateUser(user user.UserDto) (*user.UserDto, error) {
	return nil, errors.New("Some error")
}
func (r *UserRepositoryImplError) DeleteUser(id int64) error {
	return errors.New("Some error")
}
func (r *UserRepositoryImplError) GetUsers(where string) ([]user.UserDto, error) {
	return nil, errors.New("Some error")
}

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userServiceSuccess = user.NewUserService(&UserRepositoryImplSuccess{
		db: db,
	})
	userServiceError = user.NewUserService(&UserRepositoryImplError{
		db: db,
	})
}

func TestGetUserSuccess(t *testing.T) {
	result, err := userServiceSuccess.GetUser(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user.User{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, *result)
}
func TestGetUserError(t *testing.T) {
	result, err := userServiceError.GetUser(1)
	if err == nil && result != nil {
		t.Error("Error can`t be nil here or result != nil")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestGetUserByEmailSuccess(t *testing.T) {
	result, err := userServiceSuccess.GetUserByEmail("")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user.User{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, *result)
}
func TestGetUserByEmailError(t *testing.T) {
	result, err := userServiceError.GetUserByEmail("")
	if err == nil && result != nil {
		t.Error("Error can`t be nil here or result != nil")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestRegistrationSuccess(t *testing.T) {
	result, err := userServiceSuccess.Registration(user.UserDto{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, user.User{
		Id:        1,
		FirstName: "Thaddeus",
		LastName:  "Gerhold",
		Email:     "Kendall22@gmail.com",
		Password:  "it4xxOEvrY2vua7",
	}, *result)
}
func TestRegistrationError(t *testing.T) {
	result, err := userServiceError.Registration(user.UserDto{})
	if err == nil && result != nil {
		t.Error("Error can`t be nil here or result != nil")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestGetUsersSuccess(t *testing.T) {
	result, err := userServiceSuccess.GetUsers("")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, []user.User{
		{
			Id:        1,
			FirstName: "Thaddeus",
			LastName:  "Gerhold",
			Email:     "Kendall22@gmail.com",
			Password:  "it4xxOEvrY2vua7",
		},
		{
			Id:        2,
			FirstName: "Vernon",
			LastName:  "Hand",
			Email:     "Baylee0@hotmail.com",
			Password:  "1o0kLTPRTlYKgSv",
		},
	}, result)
}
func TestGetUsersError(t *testing.T) {
	result, err := userServiceError.GetUsers("")
	if err == nil && result != nil {
		t.Error("Error can`t be nil here or result != nil")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestDeleteUserSuccess(t *testing.T) {
	err := userServiceSuccess.DeleteUser(1)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteUserError(t *testing.T) {
	err := userServiceError.DeleteUser(1)
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
