package orderController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrdersByUserId(c *gin.Context) {

	log.Debug("Order id to load: " + c.Param("id_user"))
	id, _ := strconv.Atoi(c.Param("id_user"))

	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrdersByUserId(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func GetOrders(c *gin.Context) {
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func InsertOrder(c *gin.Context) {

	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var id int
	var ok bool
	id, ok = service.OrderService.CheckStock(orderDto)

	if ok == false {
		c.JSON(http.StatusBadRequest, id)
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}
