package category

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCategories() model.Categories {
	var categories model.Categories
	Db.Find(&categories)

	log.Debug("Products: ", categories)

	return categories
}

func GetCategoryIdByName(name string) model.Category {

	var category model.Category
	Db.Where("name = ?", name).Find(&category)
	log.Debug("Category: ", category)

	return category

}
