package dto

type DetailDto struct {
	Id_detail       int     `json:"id_detail"`
	Precio_Unitario float32 `json:"precio_unitario"`
	Cantidad        float32 `json:"cantidad"`
	Total           float32 `json:"total"`
	Nombre          string  `json:"nombre"`
	Id_order        int     `json:"id_order"`
	Id_product      int     `json:"id_product"`
}

type DetailsDto []DetailDto
