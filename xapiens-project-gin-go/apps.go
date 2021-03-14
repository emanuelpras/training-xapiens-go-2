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

/* func Sentry() {
	var DsnSentry string

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(err.Error())
	} else { // tidak ada error saat ambil file .env
		DsnSentry = os.Getenv("DSN_SENTRY") // isi value dari variable DsnSentry
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: DsnSentry,
	})

	// log untuk capture error
	if _, errs := os.Open("filename.ext"); errs != nil {
		sentry.CaptureException(err)
	}

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
} */
