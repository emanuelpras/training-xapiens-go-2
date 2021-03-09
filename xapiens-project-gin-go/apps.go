package main

import (
	"xapiens-day9/config"
	"xapiens-day9/controller"
	"xapiens-day9/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controller.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// routing end point
	routing := gin.Default()

	// routing vendor
	routing.POST("/vendor", strDB.PostDataVendor)
	routing.GET("/vendorList", strDB.GetVendorList)
	routing.GET("/VendorByQuery", strDB.GetDataVendor)

	// routing employee
	routing.POST("/employee", strDB.PostDataEmployee)
	routing.GET("/employeeList", strDB.GetEmployeeList)

	// routing upload files
	routing.POST("/uploadFile", controller.UploadSingleFile)
	routing.POST("/uploadMultipleFile", controller.UploadMultipleFile)

	//routing 3rd api / http request
	routing.GET("/post", controller.HttpRequest)
	routing.Run()
}
