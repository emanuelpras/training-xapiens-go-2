package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"xapiens-day9/models"
	sentryLogger "xapiens-day9/sentry"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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

	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:4455")
	}, 10)

	pool.MaxActive = 10

	// ambil 1 koneksi
	conn := pool.Get()
	defer conn.Close()

	reply, _ := redis.Bytes(conn.Do("GET", "mahasiswa"))

	if reply != nil {
		err := json.Unmarshal(reply, &vendor)

		if err != nil {
			log.Fatalln("error ->", err.Error())
		}

		if length := len(vendor); length > 0 {
			result = gin.H{
				"message": "success dari redis",
				"data":    vendor,
				"count":   length,
			}
		}
		c.JSON(http.StatusOK, result)
	} else {
		StrDB.DB.Find(&vendor, "vendor_id = ?", c.Query("vendorID"))
		result = gin.H{
			"msg":  "success dari db",
			"data": vendor,
		}

		jd, _ := json.Marshal(vendor)

		c.JSON(http.StatusOK, result)
		conn.Do("SETEX", "mahasiswa", 20, jd)
	}

}
