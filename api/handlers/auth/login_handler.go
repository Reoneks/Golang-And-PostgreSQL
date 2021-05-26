package auth

import (
	"net/http"
	"test/user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(authService user.AuthService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var loginRequest LoginRequest
		if err := ctx.Bind(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// TODO: Add Validation
		// email -  (IsEmailValid(email) bool)
		// password - (IsPasswordValid(pass) bool) не меньше 8 символов и не больше 32
		// Если не валид return StatusBadRequest

		login, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
		}

		if login == nil {
			ctx.JSON(http.StatusBadRequest)
		}

		// TODO Check without password
		ctx.JSON(http.StatusOK, login)
	}
}
