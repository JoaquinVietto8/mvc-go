package services

import (
	categoryCliente "mvc-go/clients/category"
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategories() (dto.CategoriesDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Id_category = category.Id
		categoryDto.Name = category.Name
		categoryDto.Icon = category.Icon
		categoryDto.Cantidad = productCliente.GetCantidadProducts(category.Id)

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}
