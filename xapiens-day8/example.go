package main

import (
	"fmt"
	"xapiens-day8/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()

	//ini ngikut dari dokumentasi gin
	request.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.GET("/ping", exampleFunction) //penulisan lebih clean - function exampleFunction ada di bawah function main

	// method Get untuk ambil data berdasarkan value dari parameter
	request.GET("/pings/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "pong " + name,
		})
	})

	// method Get untuk ambil data dari query
	request.GET("/ping/query/", func(c *gin.Context) {
		names := c.Query("names") // ini code yg membedakan query & parameter
		age := c.DefaultQuery("age", "17")
		c.JSON(200, gin.H{
			"message": "pong " + names,
			"äge":     age,
		})
	})

	// method POST
	// method Post bisa dapat data dari form data atau raw
	// kalau mau pakai raw, harus membuat struct terlebih dahulu.

	// code dibawah adalah method post menggunakan form data
	request.POST("/user/", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		city := c.PostForm("city")
		role := c.PostForm("role")

		c.JSON(200, gin.H{
			"name": name,
			"äge":  age,
			"city": city,
			"role": role,
		})
	})

	// method Post menggunakan raw data.
	// pertama yang harus dilakukan adalah membuat struct terlebih dahulu
	type mahasiswa struct {
		Nama        string `json:"fullName"`
		Nim         string `json:"nim"`
		Jurusan     string `json:"jurusan"`
		Fakultas    string `json:"fakultas"`
		Universitas string `json:"univ"`
	}

	// setelah membuat struct, buat routingnya
	request.POST("/mahasiswa/", func(c *gin.Context) {
		var mhs mahasiswa   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
		err := c.Bind(&mhs) // buat variable yg fungsinya untuk check error
		if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 74)
			fmt.Println("Terjadi error")
		}

		// response
		c.JSON(200, gin.H{
			"response": "success",
			"msg":      "data successfully added.",
			"data":     mhs,
			// "nama":        mhs.Nama,
			// "nim":         mhs.Nim,
			// "jurusan":     mhs.Jurusan,
			// "fakultas":    mhs.Fakultas,
			// "universitas": mhs.Universitas,
		})
	})

	// response : menampilkan data & msg nya -> implementasi map
	request.POST("/mahasiswas/", controllers.PostDataMahasiswa)

	request.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/* func exampleFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
} */
