package product_handler

import (
	"net/http"
	"test/product"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateCommentRequest struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	ProductID int64     `json:"product_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UpdateCommentHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateCommentRequest UpdateCommentRequest
		if err := ctx.BindJSON(&updateCommentRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		comment, err := productService.UpdateComment(product.CommentsDto(updateCommentRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else if comment == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't update comment",
			})
		} else {
			ctx.JSON(http.StatusOK, comment)
		}
	}
}
