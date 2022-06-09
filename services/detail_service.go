package services

import (
	detailCliente "mvc-go/clients/detail"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type detailService struct{}

type detailServiceInterface interface {
	GetDetailsByOrderId(id int) (dto.DetailsDto, e.ApiError)
}

var (
	DetailService detailServiceInterface
)

func init() {
	DetailService = &detailService{}
}

func (s *detailService) GetDetailsByOrderId(id int) (dto.DetailsDto, e.ApiError) {

	var details model.Details = detailCliente.GetDetailsByOrderId(id)
	var detailsDto dto.DetailsDto

	for _, detail := range details {
		var detailDto dto.DetailDto

		detailDto.Nombre = detail.Nombre
		detailDto.Cantidad = detail.Cantidad
		detailDto.Precio_Unitario = detail.Precio_Unitario
		detailDto.Total = detail.Total

		detailsDto = append(detailsDto, detailDto)
	}

	return detailsDto, nil
}
