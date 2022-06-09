package model

type Detail struct {
	Id              int     `gorm:"primaryKey"`
	Precio_Unitario float32 `gorm:"type:varchar(150);not null"`
	Cantidad        float32 `gorm:"type:varchar(150);not null"`
	Total           float32 `gorm:"type:varchar(150);not null"`
	Nombre          string  `gorm:"type:varchar(150);not null"`
	Id_order        int     `gorm:"type:varchar(150);not null"`
	Id_product      int     `gorm:"type:varchar(150);not null"`
}

type Details []Detail
