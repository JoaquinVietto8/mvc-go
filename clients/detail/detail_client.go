package detail

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetDetailsByOrderId(id int) model.Details {
	var details model.Details
	Db.Where("id_order = ?", id).Find(&details)
	log.Debug("Details: ", details)

	return details
}

func InsertDetails(details model.Details) model.Details {

	for _, detail := range details {

		result := Db.Create(&detail)

		if result.Error != nil {
			//TODO Manage Errors
			log.Error("")
		}
	}

	return details
}
