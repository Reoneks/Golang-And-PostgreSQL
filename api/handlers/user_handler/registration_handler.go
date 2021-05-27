package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationRequest struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min=8,max=32,alphanum"`
}

func RegistrationHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var registrationRequest RegistrationRequest
		if err := ctx.BindJSON(&registrationRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&registrationRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := userService.Registration(user.UserDto(registrationRequest))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if user == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't find user",
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}