package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type GetProductsRequest struct {
	Name *string `form:"name"`
}

func GetProductsHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getProductsRequest GetProductsRequest
		if err := ctx.Bind(&getProductsRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		filter := &product.ProductFilter{
			Name: getProductsRequest.Name,
		}
		product, err := productService.GetProducts(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, product)
		}
	}
}
