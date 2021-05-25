package product

type UserProductDto struct {
	UserID    int64 `gorm:"column:user_id"`
	ProductID int64 `gorm:"column:product_id"`
}

func (UserProductDto) TableName() string {
	return "users_products"
}
