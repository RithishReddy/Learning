package main

import (
	db "Exercise/Gorm/db"
	"flag"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct {
	gorm.Model
	Name         string
	Designation  string
	DepartmentID uint
	Department   Department
	ManagerID    *uint
	Manager      *Employee
}

type Department struct {
	gorm.Model
	Name      string
	Employees []Employee
}

func main() {

	username := flag.String("user_name", "postgres", "UserName String")
	password := flag.String("password", "", "Password string")
	database := flag.String("db_name", "employees", "databaseName String")

	flag.Parse()

	dbase := db.NewClient(*username, *password, *database)

	defer dbase.Close()

	db.SeedDB(dbase)

	db.ShowEmployeeDetails(dbase)

}
