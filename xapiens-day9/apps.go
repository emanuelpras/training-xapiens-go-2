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

	routing := gin.Default()
	//routing

	routing.POST("/vendor", strDB.PostDataVendor)
	routing.GET("/vendorList", strDB.GetVendorList)
	routing.GET("/VendorByQuery", strDB.GetDataVendor)
	routing.Run()
}
