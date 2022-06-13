package services

import (
	detailCliente "mvc-go/clients/detail"
	orderCliente "mvc-go/clients/order"
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	CheckStock(orderDto dto.OrderDto) (int, bool)
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

	var total float32
	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id_order = order.Id
		orderDto.Id_user = order.Id_user
		orderDto.Fecha = order.Fecha

		var details model.Details = detailCliente.GetDetailsByOrderId(order.Id)

		for _, detail := range details {
			var detailDto dto.DetailDto

			detailDto.Id_product = detail.Id_product

			var product model.Product = productCliente.GetProductById(detailDto.Id_product)

			detailDto.Id_detail = detail.Id
			detailDto.Precio_Unitario = product.Price
			detailDto.Cantidad = detail.Cantidad
			detailDto.Total = detail.Precio_Unitario * detail.Cantidad
			detailDto.Nombre = product.Name
			detailDto.Id_order = order.Id

			total += detailDto.Total

			orderDto.Detail = append(orderDto.Detail, detailDto)
		}

		orderDto.Total = total

		total = 0

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) CheckStock(orderDto dto.OrderDto) (int, bool) {

	var ok bool

	for _, detailDto := range orderDto.Detail {

		var detail model.Detail

		detail.Id_product = detailDto.Id_product

		ok = productCliente.CheckStock(detail.Id_product, int(detailDto.Cantidad))

		if ok == false {
			return detail.Id_product, ok
		}
	}
	return 0, true
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Id_user = orderDto.Id_user
	order.Fecha = time.Now()

	order = orderCliente.InsertOrder(order)

	var total float32
	var details model.Details

	for _, detailDto := range orderDto.Detail {

		var detail model.Detail

		detail.Id_product = detailDto.Id_product

		var product model.Product = productCliente.GetProductById(detail.Id_product)
		detail.Precio_Unitario = product.Price
		detail.Cantidad = detailDto.Cantidad
		detail.Total = detail.Precio_Unitario * detail.Cantidad
		detail.Nombre = product.Name
		detail.Id_order = order.Id

		total += detail.Total

		productCliente.UpdateStock(detail.Id_product, int(detailDto.Cantidad))

		details = append(details, detail)
	}

	orderCliente.UpdateTotal(total, order.Id)

	detailCliente.InsertDetails(details)

	return orderDto, nil
}
