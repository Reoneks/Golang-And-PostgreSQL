package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
)

type GetUsersRequest struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Email     string `form:"email"`
	Status    int64  `form:"status"`
}

func GetUsersHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getUsersRequest GetUsersRequest
		if err := ctx.Bind(&getUsersRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := userService.GetUsers(user.UserFilter(getUsersRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if user == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
