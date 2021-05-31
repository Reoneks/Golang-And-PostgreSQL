package product_handler

import (
	"net/http"
	"strconv"
	"test/product"

	"github.com/gin-gonic/gin"
)

func GetProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("productId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		product, comments, err := productService.GetProduct(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"product":  product,
				"comments": comments,
			})
		}
	}
}
