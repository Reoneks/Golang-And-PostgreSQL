package main

import (
	"log"
	"test/api"
	"test/auth"
	"test/product"
	"test/user"

	c "test/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Print("No .env file found")
	}

	config := c.NewConfig()

	db := config.DBClient()
	jwt := config.JWT()

	userRepository := user.NewUserRepository(db)
	productRepository := product.NewProductRepository(db)
	uPRepository := product.NewUserProductRepository(db)
	commentsRepository := product.NewCommentsRepository(db)

	userService := user.NewUserService(userRepository)
	productService := product.NewProductService(productRepository, uPRepository, commentsRepository)
	authService := auth.NewAuthService(userService, jwt)

	httpServer := api.NewHTTPServer(
		config.ServerAddress(),
		authService,
		userService,
		productService,
	)

	log.Printf("HTTP Server listening at: %v", config.ServerAddress().String())

	if err := httpServer.Start(); err != nil {
		panic(err)
	}

}
