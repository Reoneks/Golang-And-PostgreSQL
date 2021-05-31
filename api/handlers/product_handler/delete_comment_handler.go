package product_handler

import (
	"net/http"
	"test/product"

	"github.com/gin-gonic/gin"
)

type DeleteCommentRequest struct {
	CommentId int64 `form:"comment_id"`
}

func DeleteCommentHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var deleteCommentRequest DeleteCommentRequest
		if err := ctx.Bind(&deleteCommentRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := productService.DeleteComment(deleteCommentRequest.CommentId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "The comment was deleted successfully")
		}
	}
}
