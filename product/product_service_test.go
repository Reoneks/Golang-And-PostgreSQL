package product_test

import (
	"strconv"
	"test/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var productService product.ProductService

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	productRepository := product.NewProductRepository(db)
	uPRepository := product.NewUserProductRepository(db)
	commentsRepository := product.NewCommentsRepository(db)
	productService = product.NewProductService(productRepository, uPRepository, commentsRepository)
}

func TestCreateProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := productService.CreateProduct(product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 1,
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 1,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}

func TestAddComment(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := productService.AddComment(product.CommentsDto{
			Id:        i,
			Text:      "test comment",
			ProductID: i,
			CreatedBy: 1,
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.CommentsDto{
			Id:        i,
			Text:      "test comment",
			ProductID: i,
			CreatedBy: 1,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}

func TestGetProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, result1, err := productService.GetProduct(i)
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 1,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
		assert.Equal(t, product.CommentsDto{
			Id:        i,
			Text:      "test comment",
			ProductID: i,
			CreatedBy: 1,
			CreatedAt: result1[0].CreatedAt,
			UpdatedAt: result1[0].UpdatedAt,
		}, *result1[0])
	}
}
func TestGetProducts(t *testing.T) {
	result, err := productService.GetProducts("name:Test")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 5, len(result))
	for i, oneProduct := range result {
		assert.Equal(t, product.ProductDto{
			Id:        int64(i + 1),
			Name:      "Test" + strconv.FormatInt(int64(i+1), 10),
			CreatedBy: 1,
			CreatedAt: oneProduct.CreatedAt,
			UpdatedAt: oneProduct.UpdatedAt,
		}, *oneProduct)
	}
}

func TestUpdateProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		_, err := productService.UpdateProduct(product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 2,
		}, 2)
		if err == nil {
			t.Error("This user can't update a product")
			continue
		}
	}
	for i := int64(1); i < 6; i++ {
		result, err := productService.UpdateProduct(product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 2,
		}, 1)
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 2,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}

func TestAddUsers(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := productService.AddUsers(i, 1, []int64{2})
		if err == nil {
			t.Error("This user can't delete a product")
		}
	}
	for i := int64(1); i < 6; i++ {
		err := productService.AddUsers(i, 2, []int64{1})
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUpdateComment(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := productService.UpdateComment(product.CommentsDto{
			Id:        i,
			Text:      "test comment " + strconv.FormatInt(i, 10),
			ProductID: i,
			CreatedBy: 2,
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.CommentsDto{
			Id:        i,
			Text:      "test comment " + strconv.FormatInt(i, 10),
			ProductID: i,
			CreatedBy: 2,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}

func TestDeleteComment(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := productService.DeleteComment(i)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := productService.DeleteProduct(i, 1)
		if err == nil {
			t.Error("This user can't delete a product")
		}
	}
	for i := int64(1); i < 6; i++ {
		err := productService.DeleteProduct(i, 2)
		if err != nil {
			t.Error(err)
		}
	}
}
