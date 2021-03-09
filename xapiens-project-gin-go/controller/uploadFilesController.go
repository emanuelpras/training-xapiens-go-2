package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// function untuk upload single file
func UploadSingleFile(c *gin.Context) {
	file, _ := c.FormFile("upload-file")
	fmt.Println(file.Filename)
	path := "images/"

	// checking saat upload ada error atau tidak
	if err := c.SaveUploadedFile(file, path+file.Filename); err != nil { // ketika ada error
		fmt.Println(path)
		fmt.Println("Ada error ", err.Error())
	} else { // tidak ada error
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

// function untuk upload multiple file
func UploadMultipleFile(c *gin.Context) {
	form, _ := c.MultipartForm() // declare multiple file
	files := form.File["upload-files"]
	path := "images/"

	// bikin for untuk upload filesnya
	for _, file := range files {
		fmt.Println(file.Filename)
		// checking saat upload apakah ada error atau tidak
		if err := c.SaveUploadedFile(file, path+file.Filename); err != nil { // ketika ada error
			fmt.Println(path)
			fmt.Println("Ada error ", err.Error())
		} else { // tidak ada error

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		}
	}
}
