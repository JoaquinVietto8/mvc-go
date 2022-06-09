package dto

type CategoryDto struct {
	Id_category int    `json:"id_category"`
	Name        string `json:"name"`
}

type CategoriesDto []CategoryDto
