package product

import "test/user"

func FromProductDto(productDto ProductDto) Product {
	product := Product{
		Id:        productDto.Id,
		Name:      productDto.Name,
		CreatedBy: productDto.CreatedBy,
		CreatedAt: productDto.CreatedAt,
		UpdatedAt: productDto.UpdatedAt,
	}
	if productDto.User != nil {
		convUser := user.FromUserDto(*productDto.User)
		product.User = &convUser
	}
	return product
}

func FromProductDtos(productDtos []ProductDto) (products []Product) {
	for _, dto := range productDtos {
		products = append(products, FromProductDto(dto))
	}
	return
}

func ToProductDto(product Product) ProductDto {
	productDto := ProductDto{
		Id:        product.Id,
		Name:      product.Name,
		CreatedBy: product.CreatedBy,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	if product.User != nil {
		convUser := user.ToUserDto(*product.User)
		productDto.User = &convUser
	}
	return productDto
}

func ToProductDtos(products []Product) (productDtos []ProductDto) {
	for _, dto := range products {
		productDtos = append(productDtos, ToProductDto(dto))
	}
	return
}
