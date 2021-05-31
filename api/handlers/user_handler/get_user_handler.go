package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
)

type GetUserRequest struct {
	Id int64 `form:"id"`
}

func GetUserHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getUserRequest GetUserRequest
		if err := ctx.Bind(&getUserRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		obtainedUser, err := userService.GetUser(getUserRequest.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if obtainedUser == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "can't find user",
			})
		} else {
			ctx.JSON(http.StatusOK, obtainedUser)
		}
	}
}
