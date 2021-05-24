package product

import (
	"time"

	gm "gorm.io/gorm"
)

type ProductRepository interface {
	GetProduct(id int64) (*ProductDto, error)
	CreateProduct(user ProductDto) (*ProductDto, error)
	UpdateProduct(user ProductDto) (*ProductDto, error)
	DeleteProduct(id int64) error
	SelectAllProducts() ([]*ProductDto, error)
}

type ProductRepositoryImpl struct {
	db *gm.DB
}

func (r *ProductRepositoryImpl) GetProduct(id int64) (*ProductDto, error) {
	user := &ProductDto{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *ProductRepositoryImpl) CreateProduct(user ProductDto) (*ProductDto, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if user.Id == 0 {
		var lastUser *ProductDto
		if err := r.db.Last(&lastUser).Error; err != nil {
			return nil, err
		}
		user.Id = lastUser.Id + 1
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(user ProductDto) (*ProductDto, error) {
	user.UpdatedAt = time.Now()
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int64) error {
	if err := r.db.Delete(&ProductDto{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) SelectAllProducts() (users []*ProductDto, err error) {
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func NewProductRepository(db *gm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db,
	}
}
