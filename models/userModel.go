package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Cpf      int
	Name     string
	Category string
}
