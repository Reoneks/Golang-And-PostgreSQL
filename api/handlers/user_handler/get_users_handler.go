package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
)

type GetUsersRequest struct {
	Search string `json:"search"`
}

func GetUsersHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getUsersRequest GetUsersRequest
		if err := ctx.BindJSON(&getUsersRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := userService.GetUsers(getUsersRequest.Search)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if user == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't find users",
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
