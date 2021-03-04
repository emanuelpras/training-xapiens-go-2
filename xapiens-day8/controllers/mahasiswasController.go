package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type mahasiswa struct {
	Nama        string `json:"fullName"`
	Nim         string `json:"nim"`
	Jurusan     string `json:"jurusan"`
	Fakultas    string `json:"fakultas"`
	Universitas string `json:"univ"`
}

func PostDataMahasiswa(c *gin.Context) {
	var mhs mahasiswa   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
	err := c.Bind(&mhs) // baut variable yg fungsinya untuk check error
	if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 72)
		fmt.Println("Terjadi error")
	}

	// response
	c.JSON(200, gin.H{
		"msg": "data successfully added.",
		"data": map[string]interface{}{
			"nama":        mhs.Nama,
			"nim":         mhs.Nim,
			"jurusan":     mhs.Jurusan,
			"fakultas":    mhs.Fakultas,
			"universitas": mhs.Universitas,
			"dosen": []string{
				"Paulina",
				"Herman",
			},
		},
	})
}
