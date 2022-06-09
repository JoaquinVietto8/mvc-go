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
	router.GET("/home", categoryController.GetCategories)

	// Products Mapping
	router.GET("/products/:id_product", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/category/:name", productController.GetProductsByCategoryName)
	router.GET("/search-products/:search", productController.GetProductsBySearch)

	// Details Mapping
	router.GET("/detail/:id_order", detailController.GetDetailsByOrderId)

	// Orders Mapping
	router.GET("/order/:id_user", orderController.GetOrdersByUserId)
	router.GET("/order", orderController.GetOrders)
	router.POST("/new-order", orderController.InsertOrder)

	// Users Mapping
	router.GET("/user/:id_user", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.GET("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
