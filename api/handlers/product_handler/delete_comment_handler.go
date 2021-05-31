package product_handler

import (
	"net/http"
	"strconv"
	"test/product"

	"github.com/gin-gonic/gin"
)

func DeleteCommentHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		err = productService.DeleteComment(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The comment was deleted successfully")
		}
	}
}
