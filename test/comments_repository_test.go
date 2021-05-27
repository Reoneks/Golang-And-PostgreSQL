package test

import (
	"test/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var commentsRepository product.CommentsRepository

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres-db-for-tests port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	commentsRepository = product.NewCommentsRepository(db)
}

func TestCreateComment(t *testing.T) {
	_, err := product.NewProductRepository(db).CreateProduct(product.ProductDto{
		Id:        1,
		Name:      "email",
		CreatedBy: 1,
	})
	if err != nil {
		t.Error(err)
	}
	for i := int64(1); i < 6; i++ {
		result, err := commentsRepository.CreateComment(product.CommentsDto{
			Id:        i,
			Text:      "Test",
			CreatedBy: 1,
			ProductID: 1,
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.CommentsDto{
			Id:        i,
			Text:      "Test",
			CreatedBy: 1,
			ProductID: 1,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}
func TestUpdateComment(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		result, err := commentsRepository.UpdateComment(product.CommentsDto{
			Id:        i,
			Text:      "Test",
			CreatedBy: 2,
			ProductID: 1,
		})
		if err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, product.CommentsDto{
			Id:        i,
			Text:      "Test",
			CreatedBy: 2,
			ProductID: 1,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}, *result)
	}
}
func TestDeleteComment(t *testing.T) {
	for i := int64(1); i < 6; i++ {
		err := commentsRepository.DeleteComment(i)
		if err != nil {
			t.Error(err)
		}
	}
	err := product.NewProductRepository(db).DeleteProduct(1)
	if err != nil {
		t.Error(err)
	}
}
