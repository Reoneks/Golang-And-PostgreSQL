package product_handler

import (
	"net/http"
	"test/product"
	"time"

	"github.com/gin-gonic/gin"
)

type AddCommentRequest struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	ProductID int64     `json:"product_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AddCommentHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addCommentRequest AddCommentRequest
		if err := ctx.BindJSON(&addCommentRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		comment, err := productService.AddComment(product.Comments(addCommentRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else if comment == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusCreated, comment)
		}
	}
}
