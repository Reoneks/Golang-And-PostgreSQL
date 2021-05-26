package user_test

import (
	"test/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var authService user.AuthService

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	authService = user.NewAuthService(userService)
}

func TestLogin(t *testing.T) {
	result, err := authService.Login("Admin Admin", "Last")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(
		t,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.lh6Thek-Ki9doqAGlZnAt6MBkeamzog49VqR0aMV2os",
		result.Token,
	)
}
