package dto

type OrderDto struct {
	Id_order int        `json:"id_order"`
	Total    float32    `json:"total"`
	Id_user  int        `json:"id_user"`
	Detail   DetailsDto `json:"OrderDetails"`
}

type OrdersDto []OrderDto
