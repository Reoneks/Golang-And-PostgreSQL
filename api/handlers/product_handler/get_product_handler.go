package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type GetProductRequest struct {
	Id int64 `json:"id"`
}

func GetProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getProductRequest GetProductRequest
		if err := ctx.BindJSON(&getProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		product, comments, err := productService.GetProduct(getProductRequest.Id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't find product",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"product":  product,
				"comments": comments,
			})
		}
	}
}
