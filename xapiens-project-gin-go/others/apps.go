package main

import (
	"log"
	"xapiens-day9/config"
	"xapiens-day9/controller"
	"xapiens-day9/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controller.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// code untuk seeder model user
	models.SeederUser(dbPG)

	// routing end point
	routing := gin.Default()

	routing.POST("/login", strDB.MiddleWare().LoginHandler)

	routing.NoRoute(strDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		log.Println("ini c nya", c)
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := routing.Group("/auth")

	auth.Use(strDB.MiddleWare().MiddlewareFunc())
	{
		// routing vendor
		auth.POST("/vendor", strDB.PostDataVendor)
		auth.GET("/vendorList", strDB.GetVendorList)

		auth.GET("/VendorByQuery", strDB.GetDataVendor)

		// routing employee
		auth.POST("/employee", strDB.PostDataEmployee)
		auth.GET("/employeeList", strDB.GetEmployeeList)

		// routing upload files
		auth.POST("/uploadFile", controller.UploadSingleFile)
		auth.POST("/uploadMultipleFile", controller.UploadMultipleFile)

		//routing 3rd api / http request
		auth.GET("/post", controller.HttpRequest)
	}

	routing.Run()
}
