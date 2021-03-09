package models

import (
	"time"

	"gorm.io/gorm"
)

type Vendor struct {
	VendorID   string     `gorm:"primarykey" json:"vendorID"`
	VendorName string     `json:"vendorName"`
	ZipCode    string     `json:"zipCode"`
	Employee   []Employee `gorm:"foreignKey:VendorID" json:"employee"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
