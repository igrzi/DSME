package models

import "gorm.io/gorm"

type Spots struct {
	gorm.Model
	Quantity int
}
