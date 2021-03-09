package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID string   `gorm:"primarykey" json:"employeeID"`
	FullName   string   `json:"fullName"`
	Position   string   `json:"position"`
	VendorID   string   `json:"vendorID"`
	Vendor     []Vendor `gorm:"foreignKey:VendorID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
