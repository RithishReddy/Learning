package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func NewClient(username string, password string, dbname string) *gorm.DB {

	db, err := gorm.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", username, password, dbname))
	if err != nil {
		panic(err.Error())
	}
	return db
}
