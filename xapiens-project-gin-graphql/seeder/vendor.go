package seeder

import (
	"xapiens-project-gin-graphql/models"

	"gorm.io/gorm"
)

func VendorSeeder(db *gorm.DB) {
	var vendorArray = [...][3]string{
		{"2021-vendor-A", "Vendor ABC", "5512"},
		{"2021-vendor-B", "Vendor BCD", "5512"},
		{"2021-vendor-C", "Vendor CDE", "5512"},
		{"2021-vendor-D", "Vendor DEF", "5512"},
	}

	var vendor models.Vendor

	for _, v := range vendorArray {
		vendor.VendorID = v[0]
		vendor.VendorName = v[1]
		vendor.ZipCode = v[2]

		db.Create(&vendor)
	}
}
