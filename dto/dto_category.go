package dto

type CategoryDto struct {
	Id_category int    `json:"id_category"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Cantidad    int    `json:"cantidad"`
}

type CategoriesDto []CategoryDto
