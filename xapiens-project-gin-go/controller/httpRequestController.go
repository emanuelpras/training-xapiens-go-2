package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpRequest(c *gin.Context) {
	// ambil data dari api
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	// check kondisi saat ambil data dari api error / tidak
	if err != nil { // ini kalau error
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	// check kondisi saat membaca data body dari api
	if err != nil { // ini kalau error
		fmt.Println(err.Error())
	}

	type DataApi struct {
		UserId uint   `json:"user_id"`
		ID     uint   `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	fmt.Println(body)
	fmt.Println("==============================")
	fmt.Println(string(body)) // convert body yang valuenya byte menjadi string
	// sb := string(body)        // convert body yang valuenya byte menjadi string
	data := DataApi{}

	// unmarshal dari hasil get api
	// unmarshall itu adalah convert dari tipe data .... ke bentuk JSON / object
	if error := json.Unmarshal(body, &data); error != nil {
		fmt.Println("Error ", err.Error())
	}

	c.JSON(http.StatusOK, data)
}
