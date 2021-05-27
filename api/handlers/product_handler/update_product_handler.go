package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type UpdateProductRequest struct {
	ProductInfo product.ProductDto `json:"product_info"`
	UserId      int64              `json:"user_id"`
}

func UpdateProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateProductRequest UpdateProductRequest
		if err := ctx.BindJSON(&updateProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		product, err := productService.UpdateProduct(updateProductRequest.ProductInfo, updateProductRequest.UserId)
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
