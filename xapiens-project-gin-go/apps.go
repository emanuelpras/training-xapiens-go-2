package main

import (
	"log"
	"xapiens-day9/config"
	"xapiens-day9/controller"
	middleware "xapiens-day9/midlleware"
	"xapiens-day9/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controller.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// routing end point
	routing := gin.Default()

	// routing.POST("/login", authMiddleware.LoginHandler)
	routing.POST("/login", middleware.MiddleWare().LoginHandler)

	routing.NoRoute(middleware.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := routing.Group("/auth")

	auth.Use(middleware.MiddleWare().MiddlewareFunc())
	{
		// auth.GET("/hello", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"text": "Berhasil masuk ke JWT",
		// 	})
		// })

		auth.GET("/vendorList", strDB.GetVendorList)

		// routing vendor
		auth.POST("/vendor", strDB.PostDataVendor)

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
