package product

import "time"

type ProductDto struct {
	Id        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Product struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromProductDto(productDto ProductDto) Product {
	return Product(productDto)
}

func FromProductDtos(productDtos []ProductDto) (products []Product) {
	for _, dto := range productDtos {
		products = append(products, Product(dto))
	}
	return
}

func ToProductDto(product Product) ProductDto {
	return ProductDto(product)
}

func ToProductDtos(products []Product) (productDtos []ProductDto) {
	for _, dto := range products {
		productDtos = append(productDtos, ProductDto(dto))
	}
	return
}

func (ProductDto) TableName() string {
	return "products"
}
