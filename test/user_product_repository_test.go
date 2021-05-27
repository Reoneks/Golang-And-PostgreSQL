package test

import (
	"test/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var userProductRepository product.UserProductRepository

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	userProductRepository = product.NewUserProductRepository(db)
	productRepository = product.NewProductRepository(db)
}
func TestUserProductConnection(t *testing.T) {
	_, err := productRepository.CreateProduct(product.ProductDto{
		Id:        15,
		Name:      "email",
		CreatedBy: 1,
	})
	if err != nil {
		t.Error(err)
	}
	_, err = productRepository.CreateProduct(product.ProductDto{
		Id:        16,
		Name:      "email",
		CreatedBy: 1,
	})
	if err != nil {
		t.Error(err)
	}
	result, err := userProductRepository.CreateUserProductConnection(product.UserProductDto{
		UserID:    1,
		ProductID: 15,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.UserProductDto{
		UserID:    1,
		ProductID: 15,
	}, *result)
	result, err = userProductRepository.CreateUserProductConnection(product.UserProductDto{
		UserID:    1,
		ProductID: 16,
	})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.UserProductDto{
		UserID:    1,
		ProductID: 16,
	}, *result)
}
func TestGetProductsByUserId(t *testing.T) {
	result, err := userProductRepository.GetProductsByUserId(1)
	if err != nil {
		t.Error(err)
		t.SkipNow()
	}
	if len(result) == 0 {
		t.Error("length of result is 0")
		t.SkipNow()
	}
	assert.Equal(t, []product.ProductDto{
		{
			Id:        15,
			Name:      "email",
			CreatedBy: 1,
			CreatedAt: result[0].CreatedAt,
			UpdatedAt: result[0].UpdatedAt,
		},
		{
			Id:        16,
			Name:      "email",
			CreatedBy: 1,
			CreatedAt: result[1].CreatedAt,
			UpdatedAt: result[1].UpdatedAt,
		},
	}, result)
	err = productRepository.DeleteProduct(15)
	if err != nil {
		t.Error(err)
	}
	err = productRepository.DeleteProduct(16)
	if err != nil {
		t.Error(err)
	}
}
