package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type AddUsersRequest struct {
	ProductId  int64   `json:"product_id"`
	UserId     int64   `json:"user_id"`
	UsersToAdd []int64 `json:"other_users"`
}

func AddUsersHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addUsersRequest AddUsersRequest
		if err := ctx.BindJSON(&addUsersRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := productService.AddUsers(
			addUsersRequest.ProductId,
			addUsersRequest.UserId,
			addUsersRequest.UsersToAdd,
		)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusOK, "Users are successfully added")
		}
	}
}
