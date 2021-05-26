package main

import (
	"fmt"
	"log"
	"test/product"
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
	productRepository := product.NewProductRepository(db)
	uPRepository := product.NewUserProductRepository(db)
	commentsRepository := product.NewCommentsRepository(db)
	userService := user.NewUserService(userRepository)
	productService := product.NewProductService(productRepository, uPRepository, commentsRepository)

	errors := productService.AddUsers(1, 4, []int64{1})
	for _, err := range errors {
		log.Println(err)
	}

	authService := user.NewAuthService(userService)
	login, err := authService.Login("Last Last", "Last")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(*login)
	}

	result, err := userService.GetUsers("first_name:a")
	if err != nil {
		log.Println(err)
	} else {
		for _, oneUser := range result {
			fmt.Println(*oneUser)
		}
	}
}
