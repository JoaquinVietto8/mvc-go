package services

import (
	categoryCliente "mvc-go/clients/category"
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByCategoryName(name string) (dto.ProductsDto, e.ApiError)
	GetProductsBySearch(search string) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}

	productDto.Id_product = product.Id
	productDto.Name = product.Name
	productDto.Price = product.Price
	productDto.Picture = product.Picture
	productDto.Id_category = product.Id_category
	return productDto, nil
}

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id_product = product.Id
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Picture = product.Picture
		productDto.Id_category = product.Id_category
		productDto.Stock = product.Stock

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productService) GetProductsByCategoryName(name string) (dto.ProductsDto, e.ApiError) {

	var category model.Category = categoryCliente.GetCategoryIdByName(name)
	var products model.Products = productCliente.GetProductsByCategoryId(category.Id)

	var productsDto dto.ProductsDto

	if products == nil {
		return productsDto, e.NewBadRequestApiError("order not found")
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id_product = product.Id
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Picture = product.Picture
		productDto.Id_category = product.Id_category

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productService) GetProductsBySearch(search string) (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProductsBySearch(search)

	var productsDto dto.ProductsDto

	if products == nil {
		return productsDto, e.NewBadRequestApiError("order not found")
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Id_product = product.Id
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Picture = product.Picture
		productDto.Id_category = product.Id_category

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}
