package product

import (
	"errors"
	"strings"
)

type ProductService interface {
	GetProduct(id int64) (*ProductDto, []*CommentsDto, error)
	GetProducts(search string) ([]*ProductDto, error)
	CreateProduct(product ProductDto) (*ProductDto, error)
	DeleteProduct(product_id, userId int64) error
	UpdateProduct(product ProductDto, userId int64) (*ProductDto, error)
	AddUsers(productId, userId int64, users []int64) (errorsArray []error)
	AddComment(comment CommentsDto) (*CommentsDto, error)
	UpdateComment(comment CommentsDto) (*CommentsDto, error)
	DeleteComment(commentId int64) error
}

type ProductServiceImpl struct {
	productRepository  ProductRepository
	uPRepository       UserProductRepository
	commentsRepository CommentsRepository
}

func (s *ProductServiceImpl) GetProduct(id int64) (*ProductDto, []*CommentsDto, error) {
	return s.productRepository.GetProduct(id)
}

func (s *ProductServiceImpl) GetProducts(search string) ([]*ProductDto, error) {
	params := strings.Split(search, ":")
	if len(params) == 2 {
		return s.productRepository.GetProducts(params[0] + " LIKE '%" + params[1] + "%'")
	}
	return s.productRepository.GetProducts("")
}

func (s *ProductServiceImpl) CreateProduct(product ProductDto) (*ProductDto, error) {
	return s.productRepository.CreateProduct(product)
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

func (s *ProductServiceImpl) UpdateProduct(product ProductDto, userId int64) (*ProductDto, error) {
	result, _, err := s.GetProduct(product.Id)
	if err != nil {
		return nil, err
	} else if result.CreatedBy != userId {
		return nil, errors.New("you are not allowed to do it")
	}
	return s.productRepository.UpdateProduct(product)
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

func (s *ProductServiceImpl) AddComment(comment CommentsDto) (*CommentsDto, error) {
	return s.commentsRepository.CreateComment(comment)
}

func (s *ProductServiceImpl) UpdateComment(comment CommentsDto) (*CommentsDto, error) {
	return s.commentsRepository.UpdateComment(comment)
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
