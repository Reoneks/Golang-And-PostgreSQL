package product_handler

import (
	"net/http"
	"test/product"
	"test/user"

	"github.com/gin-gonic/gin"
)

type DeleteProductRequest struct {
	ProductId int64 `form:"product_id"`
}

func DeleteProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var deleteProductRequest DeleteProductRequest
		if err := ctx.Bind(&deleteProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		thisUser, _ := ctx.Get("user")
		err := productService.DeleteProduct(deleteProductRequest.ProductId, thisUser.(*user.User).Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The product was deleted successfully")
		}
	}
}
