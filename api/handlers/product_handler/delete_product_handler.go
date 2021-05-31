package product_handler

import (
	"net/http"
	"strconv"
	"test/product"
	"test/user"

	"github.com/gin-gonic/gin"
)

func DeleteProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("productId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		thisUser, _ := ctx.Get("user")
		err = productService.DeleteProduct(int64(id), thisUser.(*user.User).Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The product was deleted successfully")
		}
	}
}
