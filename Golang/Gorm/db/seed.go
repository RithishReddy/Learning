
package db

import "github.com/jinzhu/gorm"

func SeedDB(db *gorm.DB) {
	db.Debug().DropTableIfExists(&Employee{}, &Department{})
	db.Debug().CreateTable(&Employee{}, &Department{})

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