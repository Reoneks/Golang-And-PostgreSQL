package api

import (
	"net/http"
	"net/url"

	"test/api/handlers/auth_handler"
	"test/api/handlers/product_handler"
	"test/api/handlers/user_handler"
	"test/auth"
	"test/middleware"
	"test/product"
	"test/user"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	url            *url.URL
	authService    auth.AuthService
	userService    user.UserService
	productService product.ProductService
}

func NewHTTPServer(
	url *url.URL,
	authService auth.AuthService,
	userService user.UserService,
	productService product.ProductService,
) HTTPServer {
	return &httpServer{
		url,
		authService,
		userService,
		productService,
	}
}

func (s *httpServer) Start() error {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	//^ Auth Handlers
	router.POST("/login", auth_handler.LoginHandler(s.authService))
	router.POST("/registration", user_handler.RegistrationHandler(s.userService))

	private := router.Group("/")

	private.Use(middleware.AuthMiddleware(s.userService))
	{
		//^ User Handlers
		private.GET("/get/user", user_handler.GetUserHandler(s.userService))
		private.GET("/get/user/email", user_handler.GetUserByEmailHandler(s.userService))
		private.GET("/get/users", user_handler.GetUsersHandler(s.userService))
		private.DELETE("/delete/user", user_handler.DeleteUserHandler(s.userService))

		//^ Product Handlers
		private.GET("/get/product", product_handler.GetProductHandler(s.productService))
		private.GET("/get/products", product_handler.GetProductsHandler(s.productService))
		private.POST("/create/product", product_handler.CreateProductHandler(s.productService))
		private.DELETE("/delete/product", product_handler.DeleteProductHandler(s.productService))
		private.PUT("/update/product", product_handler.UpdateProductHandler(s.productService))
		private.POST("/add/users", product_handler.AddUsersHandler(s.productService))
		private.POST("/add/comment", product_handler.AddCommentHandler(s.productService))
		private.PUT("/update/comment", product_handler.UpdateCommentHandler(s.productService))
		private.DELETE("/delete/comment", product_handler.DeleteCommentHandler(s.productService))
	}

	server := &http.Server{
		Addr:           s.url.Host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
