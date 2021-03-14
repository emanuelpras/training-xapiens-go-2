package controller

import (
	"fmt"
	"net/http"
	"xapiens-day9/models"
	sentryLogger "xapiens-day9/sentry"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk insert data vendor ke dalam table
func (StrDB *StrDB) PostDataVendor(c *gin.Context) {
	var (
		vendor models.Vendor
		result gin.H
	)

	if err := c.Bind(&vendor); err != nil {
		fmt.Println("Terjadi kesalahan")
	} else {
		// if condition ketika ingin insert data
		if res := StrDB.DB.Create(&vendor); res.Error != nil { // kondisi ketika ada error
			// error output
			err := res.Error

			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})

			sentryLogger.Sentry(err) // push log error ke sentry

		} else { // ketika tidak terjadi error
			result = gin.H{
				"msg": "success",
				"data": map[string]interface{}{
					"vendorID":   vendor.VendorID,
					"vendorName": vendor.VendorName,
					"zipCode":    vendor.ZipCode,
				},
			}
			c.JSON(http.StatusOK, result)
		}
	}
}

// Fungsi untuk mengambil semua data vendor
func (StrDB *StrDB) GetVendorList(c *gin.Context) {
	var (
		vendor []models.Vendor
		result gin.H
	)

	StrDB.DB.
		Preload("Employee").
		Find(&vendor)
	result = gin.H{
		"msg":  "success",
		"data": vendor,
	}

	c.JSON(http.StatusOK, result)
}

// Fungsi untuk mengambil data berdasarkan parameter
func (StrDB *StrDB) GetDataVendor(c *gin.Context) {
	var (
		vendor []models.Vendor
		result gin.H
	)

	StrDB.DB.Find(&vendor, "vendor_id = ?", c.Query("vendorID"))
	result = gin.H{
		"msg":  "success",
		"data": vendor,
	}

	c.JSON(http.StatusOK, result)
}
