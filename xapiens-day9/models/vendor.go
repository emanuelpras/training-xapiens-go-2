package models

import "gorm.io/gorm"

type Vendor struct {
	VendorID   string `gorm:"primarykey" json:"vendorID"`
	VendorName string `json:"vendorName"`
	ZipCode    string `json:"zipCode"`
	gorm.Model
}
