package main

import (
	"fmt"
	"log"

	//"test/product"
	"test/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=0.0.0.0 user=postgres password=postgres dbname=postgres-test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewUserRepository(db)
	//productRepository := product.NewProductRepository(db)
	userProductRepository := user.NewUserProductRepository(db)

	userRepository.CreateUser(user.UserDto{
		Id:        1,
		FirstName: "Test",
		LastName:  "Some",
		Email:     "a@a.a",
		Password:  "none",
	})

	getUser, err := userRepository.GetUser(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(getUser)

	/*userRepository.UpdateUser(user.UserDto{
		Id:        2,
		FirstName: "James!!",
		LastName:  "I_dont_know",
		Email:     "none@none.ua",
		Password:  "Its_real?",
	})

	userRepository.DeleteUser(3)

	userRepository.CreateUser(user.UserDto{
		Id:        4,
		FirstName: "Vadim",
		LastName:  "Linets",
		Email:     "vadym.linets@computools.com",
		Password:  "nill",
	})

	productRepository.CreateProduct(product.ProductDto{
		Id:   1,
		Name: "Test1",
	})
	productRepository.CreateProduct(product.ProductDto{
		Name: "Test2",
	})
	productRepository.CreateProduct(product.ProductDto{
		Name: "Test3",
	})

	userProductRepository.CreateUserProductConnection(user.UserProductDto{
		UserID:    4,
		ProductID: 1,
	})
	userProductRepository.CreateUserProductConnection(user.UserProductDto{
		UserID:    4,
		ProductID: 2,
	})
	userProductRepository.CreateUserProductConnection(user.UserProductDto{
		UserID:    1,
		ProductID: 3,
	})*/

	getAllUsers, err := userRepository.SelectAllUsers()
	if err != nil {
		log.Println(err)
	}
	for _, user := range getAllUsers {
		fmt.Println(user)
	}

	getAllUserProducts, err := userProductRepository.GetProductsByUserId(4)
	if err != nil {
		log.Println(err)
	}
	for _, user := range getAllUserProducts {
		fmt.Println(user)
	}
}
