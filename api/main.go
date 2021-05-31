package api

import (
	"net/http"
	"net/url"

	"test/api/handlers/auth_handler"
	"test/api/handlers/product_handler"
	"test/api/handlers/user_handler"
	"test/api/middleware"
	"test/auth"
	"test/product"
	"test/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	url            *url.URL
	authService    auth.AuthService
	userService    user.UserService
	productService product.ProductService
	jwt            *jwtauth.JWTAuth
	log            *logrus.Entry
}

func NewHTTPServer(
	url *url.URL,
	authService auth.AuthService,
	userService user.UserService,
	productService product.ProductService,
	jwt *jwtauth.JWTAuth,
	log *logrus.Entry,
) HTTPServer {
	return &httpServer{
		url,
		authService,
		userService,
		productService,
		jwt,
		log,
	}
}

func (s *httpServer) Start() error {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	//^ Auth Handlers
	router.POST("/login", auth_handler.LoginHandler(s.authService))
	router.POST("/registration", user_handler.RegistrationHandler(s.userService))

	private := router.Group("/")

	private.Use(middleware.AuthMiddleware(s.userService, s.jwt))
	{
		//^ User Handlers
		private.GET("/user/:userId", user_handler.GetUserHandler(s.userService))
		private.GET("/user/email", user_handler.GetUserByEmailHandler(s.userService))
		private.GET("/users", user_handler.GetUsersHandler(s.userService))
		private.DELETE("/user", user_handler.DeleteUserHandler(s.userService))

		//^ Product Handlers
		private.GET("/product/:productId", product_handler.GetProductHandler(s.productService))
		private.GET("/products", product_handler.GetProductsHandler(s.productService))
		private.POST("/product", product_handler.CreateProductHandler(s.productService))
		private.DELETE("/product/:productId", product_handler.DeleteProductHandler(s.productService))
		private.PUT("/product", product_handler.UpdateProductHandler(s.productService))
		private.POST("/users", product_handler.AddUsersHandler(s.productService))
		private.POST("/comment", product_handler.AddCommentHandler(s.productService))
		private.PUT("/comment", product_handler.UpdateCommentHandler(s.productService))
		private.DELETE("/comment/:commentId", product_handler.DeleteCommentHandler(s.productService))
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
