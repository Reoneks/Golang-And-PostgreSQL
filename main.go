package main

import (
	"log"
	"test/api"
	"test/product"
	"test/user"

	_ "github.com/lib/pq"
	. "test/config"
)

func main() {
	_ = godotenv.Load()

	config := NewConfig()

	db := config.DBClient()
	jwt := config.JWT()

	userRepository := user.NewUserRepository(db)
	productRepository := product.NewProductRepository(db)
	uPRepository := product.NewUserProductRepository(db)
	commentsRepository := product.NewCommentsRepository(db)

	userService := user.NewUserService(userRepository)
	productService := product.NewProductService(productRepository, uPRepository, commentsRepository)
	authService := user.NewAuthService(userService, jwt)

	httpServer := api.NewHTTPServer(
		config.ServerAddress(),
		authService,
		userService,
		productService,
	)

	///  Addr http://0.0.0.0:8080
	log.Printf("HTTP Server listening at: %v", config.Addr)

	if err = httpServer.Start(); err != nil {
		panic(err)
	}

}
