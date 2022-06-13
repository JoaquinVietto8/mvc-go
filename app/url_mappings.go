package app

import (
	categoryController "mvc-go/controllers/category"
	detailController "mvc-go/controllers/detail"
	orderController "mvc-go/controllers/order"
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Category Mapping
	router.GET("/categorias", categoryController.GetCategories)

	// Products Mapping
	router.GET("/productos/:id_product", productController.GetProductById)
	router.GET("/productos", productController.GetProducts)
	router.GET("/categorias/:name", productController.GetProductsByCategoryName)
	router.GET("/search-producto/:search", productController.GetProductsBySearch)

	// Details Mapping
	router.GET("/detail/:id_order", detailController.GetDetailsByOrderId)

	// Orders Mapping
	router.GET("/orders/:id_user", orderController.GetOrdersByUserId)
	router.GET("/order", orderController.GetOrders)
	router.POST("/new-order", orderController.InsertOrder)

	// Users Mapping
	router.GET("/user/:id_user", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.GET("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
