package db

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

func ShowEmployeeDetails(db *gorm.DB) {
	employees := []Employee{}
	db.Debug().Preload("Manager").Preload("Department").Find(&employees)

	for _, employee := range employees {

		if employee.Manager != nil {
			fmt.Printf("Id: %v, Name: %v, Role:%v,Manager: %v  \n", employee.ID, employee.Name, employee.Designation, employee.Manager.Name)
		} else {
			fmt.Printf("Id: %v, Name: %v, Role:%v,Manager: Nil \n", employee.ID, employee.Name, employee.Designation)
		}
	}

}
func (e *Employee) AfterCreate() error {
	fmt.Printf("Mail sent to User %v \n", e.Name)
	return nil
}
