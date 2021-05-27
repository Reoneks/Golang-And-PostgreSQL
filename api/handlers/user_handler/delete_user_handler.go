package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
)

type DeleteUserRequest struct {
	UserId int64 `json:"user_id"`
}

func DeleteUserHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var deleteUserRequest DeleteUserRequest
		if err := ctx.BindJSON(&deleteUserRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := userService.DeleteUser(deleteUserRequest.UserId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The user was deleted successfully")
		}
	}
}
