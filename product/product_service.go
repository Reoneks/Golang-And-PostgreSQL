package product

import (
	"errors"
	"strconv"
	"strings"
)

type ProductService interface {
	GetProduct(id int64) (*Product, []Comments, error)
	GetProducts(filter ProductFilter) ([]Product, error)
	CreateProduct(product ProductDto) (*Product, error)
	DeleteProduct(product_id, userId int64) error
	UpdateProduct(product ProductDto, userId int64) (*Product, error)
	AddUsers(productId, userId int64, users []int64) (errorsArray []error)
	AddComment(comment CommentsDto) (*Comments, error)
	UpdateComment(comment CommentsDto) (*Comments, error)
	DeleteComment(commentId int64) error
}

type ProductServiceImpl struct {
	productRepository  ProductRepository
	uPRepository       UserProductRepository
	commentsRepository CommentsRepository
}

func (s *ProductServiceImpl) GetProduct(id int64) (*Product, []Comments, error) {
	result, comments, err := s.productRepository.GetProduct(id)
	if err != nil {
		return nil, nil, err
	}
	resultProduct := FromProductDto(*result)
	resultComments := FromCommentsDtos(comments)
	return &resultProduct, resultComments, nil
}

func (s *ProductServiceImpl) GetProducts(filter ProductFilter) ([]Product, error) {
	var search []string
	if filter.Name != "" {
		search = append(search, "email LIKE '%"+filter.Name+"%'")
	}
	if filter.CreatedBy != 0 {
		search = append(search, "status = '"+strconv.FormatInt(filter.CreatedBy, 10)+"'")
	}
	result, err := s.productRepository.GetProducts(strings.Join(search, " AND "))
	if err != nil {
		return nil, err
	}
	resultProducts := FromProductDtos(result)
	return resultProducts, nil
}

func (s *ProductServiceImpl) CreateProduct(product ProductDto) (*Product, error) {
	result, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	resultComments := FromProductDto(*result)
	return &resultComments, nil
}

func (s *ProductServiceImpl) DeleteProduct(product_id, userId int64) error {
	result, _, err := s.GetProduct(product_id)
	if err != nil {
		return err
	} else if result.CreatedBy != userId {
		return errors.New("you are not allowed to do it")
	}
	return s.productRepository.DeleteProduct(product_id)
}

func (s *ProductServiceImpl) UpdateProduct(product ProductDto, userId int64) (*Product, error) {
	result, _, err := s.GetProduct(product.Id)
	if err != nil {
		return nil, err
	} else if result.CreatedBy != userId {
		return nil, errors.New("you are not allowed to do it")
	}
	updateResult, err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	resultProduct := FromProductDto(*updateResult)
	return &resultProduct, nil
}

func (s *ProductServiceImpl) AddUsers(productId, userId int64, users []int64) (errorsArray []error) {
	result, _, err := s.GetProduct(productId)
	if err != nil {
		errorsArray = append(errorsArray, err)
		return
	} else if result.CreatedBy != userId {
		errorsArray = append(errorsArray, errors.New("you are not allowed to do it"))
		return
	}
	for _, user := range users {
		_, err = s.uPRepository.CreateUserProductConnection(UserProductDto{
			UserID:    user,
			ProductID: productId,
		})
		if err != nil {
			errorsArray = append(errorsArray, err)
		}
	}
	return
}

func (s *ProductServiceImpl) AddComment(comment CommentsDto) (*Comments, error) {
	result, err := s.commentsRepository.CreateComment(comment)
	if err != nil {
		return nil, err
	}
	resultComments := FromCommentsDto(*result)
	return &resultComments, nil
}

func (s *ProductServiceImpl) UpdateComment(comment CommentsDto) (*Comments, error) {
	result, err := s.commentsRepository.UpdateComment(comment)
	if err != nil {
		return nil, err
	}
	resultComments := FromCommentsDto(*result)
	return &resultComments, nil
}

func (s *ProductServiceImpl) DeleteComment(commentId int64) error {
	return s.commentsRepository.DeleteComment(commentId)
}

func NewProductService(productRepository ProductRepository, uPRepository UserProductRepository, commentsRepository CommentsRepository) ProductService {
	return &ProductServiceImpl{
		productRepository,
		uPRepository,
		commentsRepository,
	}
}
