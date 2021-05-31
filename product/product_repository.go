package product

import (
	"strings"
	"time"

	gm "gorm.io/gorm"
)

type ProductRepository interface {
	GetProduct(id int64) (*ProductDto, []CommentsDto, error)
	CreateProduct(user ProductDto) (*ProductDto, error)
	UpdateProduct(user ProductDto) (*ProductDto, error)
	DeleteProduct(id int64) error
	GetProducts(filter *ProductFilter) ([]ProductDto, error)
}

type ProductRepositoryImpl struct {
	db *gm.DB
}

func (r *ProductRepositoryImpl) GetProduct(id int64) (*ProductDto, []CommentsDto, error) {
	product := &ProductDto{}
	comments := []CommentsDto{}
	if err := r.db.Joins("User").Where("products.id = ?", id).First(product).Error; err != nil {
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

func (r *ProductRepositoryImpl) GetProducts(filter *ProductFilter) (products []ProductDto, err error) {
	var findResult *gm.DB = r.db
	var search []string
	if filter != nil {
		if filter.Name != nil {
			search = append(search, "products.name LIKE '%"+*filter.Name+"%'")
		}
		findResult = findResult.Where(strings.Join(search, " AND "))
	}
	if err := findResult.Joins("User").Find(&products).Error; err != nil {
		return nil, err
	}
	return
}

func NewProductRepository(db *gm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db,
	}
}
