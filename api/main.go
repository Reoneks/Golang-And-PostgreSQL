package api

import (
	"net/url"
	"test/api/handlers/auth"
	"test/api/http"
	"test/product"
	"test/user"
	"time"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	url            *url.URL
	authService    user.AuthService
	userService    user.UserService
	productService product.ProductService
}

func NewHTTPServer(
	url *url.URL,
	authService user.AuthService,
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
	//router.Use(apihttp.CorsMiddleware())

	private := router.Group("/")
	private.Use(apihttp.AuthMiddleware(s.authEndpoint))
	{
		private.POST("/login", auth.LoginHandler(s.authService))

		//	private.Any("/ws", apihttp.WSHandler(s.upgrader, s.connector))
		//
		//	private.GET("/rooms", s.roomController.GetRooms)
		//	private.POST("/rooms", s.roomController.Create)
		//	private.PUT("/rooms/:id", s.roomController.Update)
		//	private.DELETE("/rooms/:id", s.roomController.Delete)
		//	private.POST("/rooms/:id/users", s.roomController.AssignUsers)
		//	private.DELETE("/rooms/:id/users", s.roomController.DeleteUsers)
		//
		//	private.GET("/rooms/:id/messages", s.messageController.GetMessages)
		//	private.POST("/rooms/:id/messages/files", s.messageController.Upload)
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
