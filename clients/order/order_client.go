package order

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrdersByUserId(id int) model.Orders {
	var orders model.Orders
	log.Debug("Id: ", id)
	Db.Where("id_user = ?", id).Find(&orders)
	log.Debug("Order: ", orders)

	return orders
}

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)
	log.Debug("Orders: ", orders)

	return orders
}

func InsertOrder(order model.Order) model.Order {

	result := Db.Create(&order)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}

func UpdateTotal(total float32, id int) {

	result := Db.Model(&model.Order{}).Where("id = ?", id).Update("total", total)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("Order no encontrada")
	}

	return
}
