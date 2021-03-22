package seeder

import (
	"xapiens-project-gin-graphql/models"

	"gorm.io/gorm"
)

func EmployeSeeder(db *gorm.DB) {
	var employeeArray = [...][4]string{
		{"2021-VA-EDP", "Emanuel Dina Prasetyawan", "Trainer", "2021-vendor-A"},
		{"2021-VA-SN", "Sahlan Nasution", "Developer", "2021-vendor-B"},
		{"2021-VA-DA", "Dimas Aninditya", "Developer", "2021-vendor-C"},
	}

	var employee models.Employee

	for _, v := range employeeArray {
		employee.EmployeeID = v[0]
		employee.FullName = v[1]
		employee.Position = v[2]
		employee.VendorID = v[3]

		db.Create(&employee)
	}
}
