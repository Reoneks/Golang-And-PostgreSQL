package product

import "time"

type ProductDto struct {
	Id        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ProductDto) TableName() string {
	return "products"
}
