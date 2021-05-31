package product_handler

import (
	"net/http"
	"test/product"
	"test/user"

	"github.com/gin-gonic/gin"
)

type AddUsersRequest struct {
	ProductId  int64   `json:"product_id"`
	UsersToAdd []int64 `json:"other_users"`
}

func AddUsersHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addUsersRequest AddUsersRequest
		if err := ctx.BindJSON(&addUsersRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		thisUser, _ := ctx.Get("user")
		err := productService.AddUsers(
			addUsersRequest.ProductId,
			thisUser.(*user.User).Id,
			addUsersRequest.UsersToAdd,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}
	}
}
