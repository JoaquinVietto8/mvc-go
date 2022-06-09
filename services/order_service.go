package services

import (
	detailCliente "mvc-go/clients/detail"
	orderCliente "mvc-go/clients/order"
	prdouctCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id_order = order.Id
		orderDto.Total = order.Total
		orderDto.Id_user = order.Id_user

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrdersByUserId(id)
	var ordersDto dto.OrdersDto

	if orders == nil {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id_order = order.Id
		orderDto.Total = order.Total
		orderDto.Id_user = order.Id_user

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Total = orderDto.Total
	order.Id_user = orderDto.Id_user

	order = orderCliente.InsertOrder(order)

	var details model.Details
	var total float32

	for _, detailDto := range orderDto.Detail {

		var detail model.Detail
		detail.Id_product = detailDto.Id_product

		var product model.Product = prdouctCliente.GetProductById(detail.Id_product)
		detail.Precio_Unitario = product.Price
		detail.Cantidad = detailDto.Cantidad
		detail.Total = detail.Precio_Unitario * detail.Cantidad
		detail.Nombre = product.Name
		detail.Id_order = order.Id

		total = total + detail.Total

		details = append(details, detail)
	}

	orderCliente.UpdateTotal(total, order.Id)

	detailCliente.InsertDetails(details)

	return orderDto, nil
}
