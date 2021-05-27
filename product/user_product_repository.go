package product

import (
	"log"

	gm "gorm.io/gorm"
)

type UserProductRepository interface {
	GetProductsByUserId(id int64) ([]ProductDto, error)
	CreateUserProductConnection(user UserProductDto) (*UserProductDto, error)
}

type UserProductRepositoryImpl struct {
	db *gm.DB
}

func (r *UserProductRepositoryImpl) GetProductsByUserId(id int64) (products []ProductDto, err error) {
	var userProducts []UserProductDto
	if err = r.db.Where("user_id = ?", id).Find(&userProducts).Error; err != nil {
		return
	}
	for _, user := range userProducts {
		var productStruct *ProductDto
		if err1 := r.db.Where("id = ?", user.ProductID).First(&productStruct).Error; err1 != nil {
			log.Println(err1)
			continue
		}
		products = append(products, *productStruct)
	}
	return
}

func (r *UserProductRepositoryImpl) CreateUserProductConnection(user UserProductDto) (*UserProductDto, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserProductRepository(db *gm.DB) UserProductRepository {
	return &UserProductRepositoryImpl{
		db,
	}
}
