package product_handler

import (
	"net/http"
	"test/product"
	"test/user"

	"github.com/gin-gonic/gin"
)

type UpdateProductRequest struct {
	ProductInfo product.Product `json:"product_info"`
}

func UpdateProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateProductRequest UpdateProductRequest
		if err := ctx.BindJSON(&updateProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		thisUser, _ := ctx.Get("user")
		product, err := productService.UpdateProduct(
			product.ToProductDto(updateProductRequest.ProductInfo),
			thisUser.(*user.User).Id,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't update product",
			})
		} else {
			ctx.JSON(http.StatusOK, product)
		}
	}
}
