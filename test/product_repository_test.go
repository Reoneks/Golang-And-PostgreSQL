package test

import (
	"strconv"
	"test/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var productRepository product.ProductRepository

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	productRepository = product.NewProductRepository(db)
}

func TestCreateProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := productRepository.CreateProduct(product.ProductDto{
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

func TestGetProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, _, err := productRepository.GetProduct(i)
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
func TestGetProducts(t *testing.T) {
	result, err := productRepository.GetProducts("name LIKE '%Test%'")
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
		}, oneProduct)
	}
}

func TestUpdateProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := productRepository.UpdateProduct(product.ProductDto{
			Id:        i,
			Name:      "Test" + strconv.FormatInt(i, 10),
			CreatedBy: 2,
		})
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

func TestDeleteProduct(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := productRepository.DeleteProduct(i)
		if err != nil {
			t.Error(err)
		}
	}
}
