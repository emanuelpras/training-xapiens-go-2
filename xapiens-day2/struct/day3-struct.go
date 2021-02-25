package main

import "fmt"

type mahasiswa struct {
	nama string
	asal string
	age  int
	ipk  float64
}

func main() {

	// Deklarasi cara pertama
	var budi mahasiswa
	budi.nama = "Budi"
	budi.asal = "Sleman"
	budi.age = 18
	budi.ipk = 3.5

	fmt.Println(budi)
	fmt.Println(budi.age)
	fmt.Println("=======================================")

	// Deklarasi cara kedua
	dina := mahasiswa{
		nama: "Dina",
		asal: "Surabaya",
		age:  17,
		ipk:  3.8,
	}

	fmt.Println(dina)
	fmt.Println(dina.ipk)

	// Deklarasi cara ketiga
	// harus hati2 karena apabila parameternya kurang bisa akibatkan error
	villa := mahasiswa{"Villa", "Banjarmasin", 19, 3.4}

	fmt.Println(villa)
	fmt.Println(villa.asal)
}
