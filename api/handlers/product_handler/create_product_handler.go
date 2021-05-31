package product_handler

import (
	"net/http"
	"test/product"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateProductHandler(productService product.ProductService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var createProductRequest CreateProductRequest
		if err := ctx.BindJSON(&createProductRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		productModel := product.Product{
			Id:        createProductRequest.Id,
			Name:      createProductRequest.Name,
			CreatedBy: createProductRequest.CreatedBy,
			CreatedAt: createProductRequest.CreatedAt,
			UpdatedAt: createProductRequest.UpdatedAt,
		}
		product, err := productService.CreateProduct(productModel)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if product == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(201, product)
		}
	}
}
