package test

import (
	"errors"
	"test/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productServiceSuccess product.ProductService
var productServiceError product.ProductService

type ProductRepositoryImplSuccess struct{}

func (r *ProductRepositoryImplSuccess) GetProduct(id int64) (*product.ProductDto, []product.CommentsDto, error) {
	return &product.ProductDto{
			Id:        1,
			Name:      "Schmidt and Sons",
			CreatedBy: 1,
		}, []product.CommentsDto{
			{
				Id:        1,
				Text:      "dolor molestiae tenetur",
				ProductID: 1,
				CreatedBy: 1,
			},
			{
				Id:        2,
				Text:      "porro vel voluptas",
				ProductID: 1,
				CreatedBy: 1,
			},
		}, nil
}
func (r *ProductRepositoryImplSuccess) CreateProduct(nProduct product.ProductDto) (*product.ProductDto, error) {
	return &product.ProductDto{
		Id:        1,
		Name:      "Schmidt and Sons",
		CreatedBy: 1,
	}, nil
}
func (r *ProductRepositoryImplSuccess) UpdateProduct(nProduct product.ProductDto) (*product.ProductDto, error) {
	return &product.ProductDto{
		Id:        1,
		Name:      "Schmidt and Sons",
		CreatedBy: 1,
	}, nil
}
func (r *ProductRepositoryImplSuccess) DeleteProduct(id int64) error {
	return nil
}
func (r *ProductRepositoryImplSuccess) GetProducts(where string) ([]product.ProductDto, error) {
	return []product.ProductDto{
		{
			Id:        1,
			Name:      "Schmidt and Sons",
			CreatedBy: 1,
		},
		{
			Id:        2,
			Name:      "Zieme Inc",
			CreatedBy: 1,
		},
	}, nil
}

type CommentsRepositoryImplSuccess struct{}

func (r *CommentsRepositoryImplSuccess) CreateComment(comment product.CommentsDto) (*product.CommentsDto, error) {
	return &product.CommentsDto{
		Id:        1,
		Text:      "perferendis fugit debitis",
		ProductID: 1,
		CreatedBy: 1,
	}, nil
}
func (r *CommentsRepositoryImplSuccess) UpdateComment(comment product.CommentsDto) (*product.CommentsDto, error) {
	return &product.CommentsDto{
		Id:        1,
		Text:      "perferendis fugit debitis",
		ProductID: 1,
		CreatedBy: 1,
	}, nil
}
func (r *CommentsRepositoryImplSuccess) DeleteComment(id int64) error {
	return nil
}

type UserProductRepositoryImplSuccess struct{}

func (r *UserProductRepositoryImplSuccess) GetProductsByUserId(id int64) ([]product.ProductDto, error) {
	return []product.ProductDto{
		{
			Id:        1,
			Name:      "Schmidt and Sons",
			CreatedBy: 1,
		},
		{
			Id:        2,
			Name:      "Zieme Inc",
			CreatedBy: 1,
		},
	}, nil
}

func (r *UserProductRepositoryImplSuccess) CreateUserProductConnection(user product.UserProductDto) (*product.UserProductDto, error) {
	return &product.UserProductDto{
		UserID:    1,
		ProductID: 1,
	}, nil
}

type ProductRepositoryImplError struct{}

func (r *ProductRepositoryImplError) GetProduct(id int64) (*product.ProductDto, []product.CommentsDto, error) {
	return nil, nil, errors.New("Some error")
}
func (r *ProductRepositoryImplError) CreateProduct(nProduct product.ProductDto) (*product.ProductDto, error) {
	return nil, errors.New("Some error")
}
func (r *ProductRepositoryImplError) UpdateProduct(nProduct product.ProductDto) (*product.ProductDto, error) {
	return nil, errors.New("Some error")
}
func (r *ProductRepositoryImplError) DeleteProduct(id int64) error {
	return errors.New("Some error")
}
func (r *ProductRepositoryImplError) GetProducts(where string) ([]product.ProductDto, error) {
	return nil, errors.New("Some error")
}

type CommentsRepositoryImplError struct{}

func (r *CommentsRepositoryImplError) CreateComment(comment product.CommentsDto) (*product.CommentsDto, error) {
	return nil, errors.New("Some error")
}
func (r *CommentsRepositoryImplError) UpdateComment(comment product.CommentsDto) (*product.CommentsDto, error) {
	return nil, errors.New("Some error")
}
func (r *CommentsRepositoryImplError) DeleteComment(id int64) error {
	return errors.New("Some error")
}

type UserProductRepositoryImplError struct{}

func (r *UserProductRepositoryImplError) GetProductsByUserId(id int64) ([]product.ProductDto, error) {
	return nil, errors.New("Some error")
}
func (r *UserProductRepositoryImplError) CreateUserProductConnection(user product.UserProductDto) (*product.UserProductDto, error) {
	return nil, errors.New("Some error")
}

func init() {
	productServiceSuccess = product.NewProductService(
		&ProductRepositoryImplSuccess{},
		&UserProductRepositoryImplSuccess{},
		&CommentsRepositoryImplSuccess{},
	)
	productServiceError = product.NewProductService(
		&ProductRepositoryImplError{},
		&UserProductRepositoryImplError{},
		&CommentsRepositoryImplError{},
	)
}

func TestGetProductSuccess(t *testing.T) {
	result, comments, err := productServiceSuccess.GetProduct(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.Product{
		Id:        1,
		Name:      "Schmidt and Sons",
		CreatedBy: 1,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, *result)
	assert.Equal(t, []product.Comments{
		{
			Id:        1,
			Text:      "dolor molestiae tenetur",
			ProductID: 1,
			CreatedBy: 1,
			CreatedAt: comments[0].CreatedAt,
			UpdatedAt: comments[0].UpdatedAt,
		},
		{
			Id:        2,
			Text:      "porro vel voluptas",
			ProductID: 1,
			CreatedBy: 1,
			CreatedAt: comments[1].CreatedAt,
			UpdatedAt: comments[1].UpdatedAt,
		},
	}, comments)
}
func TestGetProductsSuccess(t *testing.T) {
	result, err := productServiceSuccess.GetProducts(product.ProductFilter{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, []product.Product{
		{
			Id:        1,
			Name:      "Schmidt and Sons",
			CreatedBy: 1,
			CreatedAt: result[0].CreatedAt,
			UpdatedAt: result[0].UpdatedAt,
		},
		{
			Id:        2,
			Name:      "Zieme Inc",
			CreatedBy: 1,
			CreatedAt: result[1].CreatedAt,
			UpdatedAt: result[1].UpdatedAt,
		},
	}, result)
}
func TestCreateProductSuccess(t *testing.T) {
	result, err := productServiceSuccess.CreateProduct(product.ProductDto{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.Product{
		Id:        1,
		Name:      "Schmidt and Sons",
		CreatedBy: 1,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, *result)
}
func TestUpdateProductSuccess(t *testing.T) {
	result, err := productServiceSuccess.UpdateProduct(product.ProductDto{}, 1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.Product{
		Id:        1,
		Name:      "Schmidt and Sons",
		CreatedBy: 1,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, *result)
}
func TestAddUsersSuccess(t *testing.T) {
	err := productServiceSuccess.AddUsers(1, 1, []int64{2, 3})
	if err != nil {
		t.Error(err)
	}
}
func TestAddCommentSuccess(t *testing.T) {
	result, err := productServiceSuccess.AddComment(product.CommentsDto{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.Comments{
		Id:        1,
		Text:      "perferendis fugit debitis",
		ProductID: 1,
		CreatedBy: 1,
	}, *result)
}
func TestUpdateCommentSuccess(t *testing.T) {
	result, err := productServiceSuccess.AddComment(product.CommentsDto{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, product.Comments{
		Id:        1,
		Text:      "perferendis fugit debitis",
		ProductID: 1,
		CreatedBy: 1,
	}, *result)
}
func TestDeleteCommentSuccess(t *testing.T) {
	err := productServiceSuccess.DeleteComment(1)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteProductSuccess(t *testing.T) {
	err := productServiceSuccess.DeleteProduct(1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetProductError(t *testing.T) {
	_, _, err := productServiceError.GetProduct(1)
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestGetProductsError(t *testing.T) {
	_, err := productServiceError.GetProducts(product.ProductFilter{})
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestCreateProductError(t *testing.T) {
	_, err := productServiceError.CreateProduct(product.ProductDto{})
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestUpdateProductError(t *testing.T) {
	_, err := productServiceError.UpdateProduct(product.ProductDto{}, 1)
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestAddUsersError(t *testing.T) {
	err := productServiceError.AddUsers(1, 1, []int64{2, 3})
	if err == nil {
		t.Error("Error can`t be nil here")
	}
}
func TestAddCommentError(t *testing.T) {
	_, err := productServiceError.AddComment(product.CommentsDto{})
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestUpdateCommentError(t *testing.T) {
	_, err := productServiceError.AddComment(product.CommentsDto{})
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestDeleteCommentError(t *testing.T) {
	err := productServiceError.DeleteComment(1)
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
func TestDeleteProductError(t *testing.T) {
	err := productServiceError.DeleteProduct(1, 1)
	if err == nil {
		t.Error("Error can`t be nil here")
	}
	assert.Equal(t, err.Error(), "Some error")
}
