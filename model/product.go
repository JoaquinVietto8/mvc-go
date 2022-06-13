package model

type Product struct {
	Id          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(350);not null"`
	Picture     string  `gorm:"type:varchar(350)"`
	Price       float32 `gorm:"type:varchar(150);not null"`
	Stock       int     `gorm:"type:varchar(150)"`
	Id_category int     `gorm:"type:varchar(150)"`
}

type Products []Product
