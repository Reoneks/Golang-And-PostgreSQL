package user_test

import (
	"strconv"
	"test/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var userService user.UserService

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepository := user.NewUserRepository(db)
	userService = user.NewUserService(userRepository)
}

func TestRegistration(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := userService.Registration(user.UserDto{
			Id:        i + 2,
			FirstName: "User" + strconv.FormatInt(i, 10),
			LastName:  "Last" + strconv.FormatInt(i, 10),
			Email:     "Sasha_Rogahn@hotmail.com",
			Password:  "uRVxm4UK61Xh4dG",
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, user.UserDto{
			Id:        i + 2,
			FirstName: "User" + strconv.FormatInt(i, 10),
			LastName:  "Last" + strconv.FormatInt(i, 10),
			Email:     "Sasha_Rogahn@hotmail.com",
			Password:  result.Password,
		}, *result)
	}
}
func TestGetUser(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := userService.GetUser(i + 2)
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, user.UserDto{
			Id:        i + 2,
			FirstName: "User" + strconv.FormatInt(i, 10),
			LastName:  "Last" + strconv.FormatInt(i, 10),
			Email:     "Sasha_Rogahn@hotmail.com",
			Password:  result.Password,
		}, *result)
	}
}
func TestGetUserByUsername(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := userService.GetUserByUsername("User" + strconv.FormatInt(i, 10) + " " + "Last" + strconv.FormatInt(i, 10))
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, user.UserDto{
			Id:        i + 2,
			FirstName: "User" + strconv.FormatInt(i, 10),
			LastName:  "Last" + strconv.FormatInt(i, 10),
			Email:     "Sasha_Rogahn@hotmail.com",
			Password:  result.Password,
		}, *result)
	}
}
func TestGetUsers(t *testing.T) {
	result, err := userService.GetUsers("email:hotmail.com")
	if err != nil {
		t.Error(err)
	} else {
		for i, userInResult := range result {
			assert.Equal(t, user.UserDto{
				Id:        int64(i) + 3,
				FirstName: "User" + strconv.FormatInt(int64(i)+1, 10),
				LastName:  "Last" + strconv.FormatInt(int64(i)+1, 10),
				Email:     "Sasha_Rogahn@hotmail.com",
				Password:  userInResult.Password,
			}, *userInResult)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := userService.DeleteUser(i + 2)
		if err != nil {
			t.Error(err)
		}
	}
}
