package product_handler

import (
	"net/http"
	"test/product"
	"test/user"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateProductRequest struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UpdateProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateProductRequest UpdateProductRequest
		if err := ctx.BindJSON(&updateProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		productModel := product.Product{
			Id:        updateProductRequest.Id,
			Name:      updateProductRequest.Name,
			CreatedBy: updateProductRequest.CreatedBy,
			CreatedAt: updateProductRequest.CreatedAt,
			UpdatedAt: updateProductRequest.UpdatedAt,
		}
		thisUser, _ := ctx.Get("user")
		product, err := productService.UpdateProduct(
			productModel,
			thisUser.(*user.User).Id,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, product)
		}
	}
}
