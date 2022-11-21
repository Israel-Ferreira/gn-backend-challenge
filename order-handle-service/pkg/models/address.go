package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City         string
	Street       string
	ZipCode      string
	Uf           string
	Neighborhood string
	Lat          float64 `gorm:"column:latitude"`
	Lng          float64 `gorm:"column:longitude"`
	OrderRefer   uint
}
