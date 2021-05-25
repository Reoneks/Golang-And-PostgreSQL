package main

import (
	"fmt"
	"log"
	"test/user"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	result, err := userService.GetUsers("first_name:a")
	if err != nil {
		log.Println(err)
	} else {
		for _, oneUser := range result {
			fmt.Println(*oneUser)
		}
	}
}
