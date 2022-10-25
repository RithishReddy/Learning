package main

import (
	"fmt"

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

	db, err := gorm.Open("postgres", "user=postgres password=Reddy@1234 dbname=employees sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	seedDB(db)

	employees := []Employee{}
	db.Debug().Preload("Manager").Preload("Department").Find(&employees)

	for _,employee:= range employees {

		if employee.Manager != nil{
			fmt.Printf("Id: %v, Name: %v, Role:%v,Manager: %v  \n",employee.ID,employee.Name,employee.Designation,employee.Manager.Name)
		}else{
			fmt.Printf("Id: %v, Name: %v, Role:%v,Manager: Nil \n",employee.ID,employee.Name,employee.Designation)
		}
	}

}

func seedDB(db *gorm.DB) {
	db.DropTableIfExists(&Employee{}, &Department{})
	db.CreateTable(&Employee{}, &Department{})

	department := []Department{
		{Name: "Engineer"},
		{Name: "Testing"},
		{Name: "SRE"},
		{Name: "Support"},
	}
	for _, val := range department {
		db.Create(&val)
	}

	role := &Department{Name: "Developer"}
	db.Create(&role)
	manager := Employee{Name: "NG", Designation: "manager", DepartmentID: role.ID}
	db.Save(&manager)

	employee := []Employee{
		{Name: "Rithish", Designation: "Developer", DepartmentID: role.ID, Manager: &manager},
		{Name: "Teja", Designation: "Developer", DepartmentID: role.ID, Manager: &manager},
	}

	for _, user := range employee {
		db.Save(&user)
	}

}


func (e *Employee) AfterCreate() error{
	fmt.Printf("Mail sent to User %v \n",e.Name)
	return nil
}