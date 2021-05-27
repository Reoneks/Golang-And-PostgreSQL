package product

import (
	"time"

	gm "gorm.io/gorm"
)

type ProductRepository interface {
	GetProduct(id int64) (*ProductDto, []CommentsDto, error)
	CreateProduct(user ProductDto) (*ProductDto, error)
	UpdateProduct(user ProductDto) (*ProductDto, error)
	DeleteProduct(id int64) error
	GetProducts(where string) ([]ProductDto, error)
}

type ProductRepositoryImpl struct {
	db *gm.DB
}

func (r *ProductRepositoryImpl) GetProduct(id int64) (*ProductDto, []CommentsDto, error) {
	product := &ProductDto{}
	comments := []CommentsDto{}
	if err := r.db.Where("id = ?", id).First(product).Error; err != nil {
		return nil, nil, err
	}
	if err := r.db.Where("product_id = ?", id).Find(&comments).Error; err != nil {
		return nil, nil, err
	}
	return product, comments, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product ProductDto) (*ProductDto, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	if product.Id == 0 {
		var lastProduct *ProductDto
		if err := r.db.Last(&lastProduct).Error; err != nil {
			return nil, err
		}
		product.Id = lastProduct.Id + 1
	}
	if err := r.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product ProductDto) (*ProductDto, error) {
	result, _, err := r.GetProduct(product.Id)
	if err != nil {
		return nil, err
	}
	product.CreatedAt = result.CreatedAt
	product.UpdatedAt = time.Now()
	if err := r.db.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int64) error {
	if err := r.db.Delete(&ProductDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) GetProducts(where string) (products []ProductDto, err error) {
	var findResult *gm.DB = r.db
	if where != "" {
		findResult = findResult.Where(where)
	}
	if err := findResult.Find(&products).Error; err != nil {
		return nil, err
	}
	return
}

func NewProductRepository(db *gm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db,
	}
}
