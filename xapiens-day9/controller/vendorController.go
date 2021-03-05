package controller

import (
	"fmt"
	"net/http"
	"xapiens-day9/models"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) PostDataVendor(c *gin.Context) {
	var (
		vendor models.Vendor
		result gin.H
	)

	if err := c.Bind(&vendor); err != nil {
		fmt.Println("Terjadi kesalahan")
	} else {
		StrDB.DB.Create(&vendor)
		result = gin.H{
			"msg": "success",
			"data": map[string]interface{}{
				"id":         vendor.ID,
				"vendorID":   vendor.VendorID,
				"vendorName": vendor.VendorName,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}

func (StrDB *StrDB) GetVendorList(c *gin.Context) {
	var (
		vendor []models.Vendor
		result gin.H
	)

	StrDB.DB.Find(&vendor)
	result = gin.H{
		"msg":  "success",
		"data": vendor,
	}

	c.JSON(http.StatusOK, result)
}

// GetDataVendor
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
