package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type GetProductsRequest struct {
	Search string `json:"search"`
}

func GetProductsHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getProductsRequest GetProductsRequest
		if err := ctx.BindJSON(&getProductsRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		product, err := productService.GetProducts(getProductsRequest.Search)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't find product",
			})
		} else {
			ctx.JSON(http.StatusOK, product)
		}
	}
}
