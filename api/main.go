package api

import (
	"net/http"
	"net/url"

	"test/api/handlers/auth_handler"
	"test/api/handlers/product_handler"
	"test/api/handlers/user_handler"
	"test/auth"
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
	//FIXME: router.Use(apihttp.CorsMiddleware())

	private := router.Group("/")
	//FIXME: private.Use(){}

	//^ Auth Handlers
	private.POST("/login", auth_handler.LoginHandler(s.authService))

	//^ User Handlers
	private.GET("/get_user", user_handler.GetUserHandler(s.userService))
	private.GET("/get_user_by_email", user_handler.GetUserByEmailHandler(s.userService))
	private.GET("/get_users", user_handler.GetUsersHandler(s.userService))
	private.POST("/registration", user_handler.RegistrationHandler(s.userService))
	private.POST("/delete_user", user_handler.DeleteUserHandler(s.userService))

	//^ Product Handlers
	private.GET("/get_product", product_handler.GetProductHandler(s.productService))
	private.GET("/get_products", product_handler.GetProductsHandler(s.productService))
	private.POST("/create_product", product_handler.CreateProductHandler(s.productService))
	private.POST("/delete_product", product_handler.DeleteProductHandler(s.productService))
	private.POST("/update_product", product_handler.UpdateProductHandler(s.productService))
	private.POST("/add_users", product_handler.AddUsersHandler(s.productService))
	private.POST("/add_comment", product_handler.AddCommentHandler(s.productService))
	private.POST("/update_comment", product_handler.UpdateCommentHandler(s.productService))
	private.POST("/delete_comment", product_handler.DeleteCommentHandler(s.productService))

	server := &http.Server{
		Addr:           s.url.Host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()

}
