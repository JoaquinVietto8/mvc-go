package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id_product"))

	id, _ := strconv.Atoi(c.Param("id_product"))
	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByCategoryName(c *gin.Context) {

	log.Debug("Category to load: " + c.Param("name"))
	name := c.Param("name")

	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProductsByCategoryName(name)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsBySearch(c *gin.Context) {

	log.Debug("Product name to load: " + c.Param("search"))
	search := c.Param("search")

	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProductsBySearch(search)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
