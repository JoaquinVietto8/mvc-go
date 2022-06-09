package model

type Order struct {
	Id      int     `gorm:"primaryKey"`
	Total   float32 `gorm:"type:varchar(150);not null"`
	Id_user int     `gorm:"type:varchar(150);not null"`
}

type Orders []Order
