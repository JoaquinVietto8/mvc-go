package model

type Category struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(350);not null"`
}

type Categories []Category
