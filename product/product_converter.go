package product

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
