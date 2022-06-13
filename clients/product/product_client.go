package product

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var product model.Product

	Db.Where("id = ?", id).First(&product)
	log.Debug("Product: ", product)

	return product
}

func CheckStock(id int, cant int) bool {
	var product model.Product
	var ok bool = true

	Db.Where("id = ?", id).First(&product)
	log.Debug("Product Stock: ", product.Stock)
	log.Debug("Cantidad solicitada: ", cant)
	if product.Stock < cant {
		ok = false
	}

	return ok
}

func UpdateStock(id int, cant int) {

	var product model.Product
	Db.Where("id = ?", id).First(&product)
	log.Debug("Product Stock: ", product.Stock)
	result := Db.Model(&model.Product{}).Where("id = ?", id).Update("stock", product.Stock-cant)
	log.Debug("Product New Stock: ", product.Stock)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("Producto no encontrado")
	}

	return
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)

	log.Debug("Products: ", products)

	return products
}

func GetProductsByCategoryId(id int) model.Products {
	var products model.Products
	Db.Where("id_category = ?", id).Find(&products)
	log.Debug("Products: ", products)

	return products
}

func GetProductsBySearch(text string) model.Products {
	var products model.Products
	Db.Where("name LIKE ?", "%"+text+"%").Find(&products)
	log.Debug("Products: ", products)

	return products
}

func GetCantidadProducts(id int) int {
	var total int
	result := Db.Raw(`SELECT COUNT(products.id) AS total
                      FROM products
                      WHERE id_category = ?`, id).Row()
	result.Scan(&total)
	log.Debug("Cant Products: ", total)

	return total
}
