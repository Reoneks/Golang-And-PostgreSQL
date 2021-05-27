package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type DeleteProductRequest struct {
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}

func DeleteProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var deleteProductRequest DeleteProductRequest
		if err := ctx.BindJSON(&deleteProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := productService.DeleteProduct(deleteProductRequest.ProductId, deleteProductRequest.UserId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The product was deleted successfully")
		}
	}
}
