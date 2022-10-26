package db

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name      string
	Employees []Employee
}