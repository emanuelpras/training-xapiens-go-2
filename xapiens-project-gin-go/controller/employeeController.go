package controller

import (
	"fmt"
	"net/http"
	"xapiens-day9/models"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk insert data employee ke dalam table
func (StrDB *StrDB) PostDataEmployee(c *gin.Context) {
	var (
		employee models.Employee
		result   gin.H
	)

	if err := c.Bind(&employee); err != nil {
		fmt.Println(err.Error())
	} else {
		StrDB.DB.Create(&employee)
		result = gin.H{
			"msg": "success",
			"data": map[string]interface{}{
				"vendorID":     employee.VendorID,
				"employeeName": employee.FullName,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}

// Fungsi untuk mengambil semua data employee
func (StrDB *StrDB) GetEmployeeList(c *gin.Context) {
	var (
		employee []models.Employee
		result   gin.H
	)

	StrDB.DB.Preload("Vendor").
		Find(&employee)
	result = gin.H{
		"msg":  "success",
		"data": employee,
	}

	c.JSON(http.StatusOK, result)
}
