package main

import "fmt"

type mahasiswa struct {
	name string
	asal string
	age  int
	ipk  float32
}

func main() {
	//Deklarasi cara pertama
	var budi mahasiswa
	budi.name = "Budi Santosa"
	budi.asal = "Jombang"
	// budi.age = 18
	budi.ipk = 3.15

	fmt.Println(budi) // hasilnya apa? budi santosa, ...., 3.15
	fmt.Println(budi.name)
	fmt.Println("==================================")

	//Deklarasi cara kedua
	dina := mahasiswa{
		name: "Elisabeth Dina",
		asal: "Surabaya",
		age:  17,
		// ipk:  3.5,
	}

	fmt.Println(dina) // hasilnya apa? budi santosa, ...., 3.15
	fmt.Println(dina.age)
	fmt.Println("==================================")

	// Deklarasi cara ketiga
	villa := mahasiswa{"Villa Delvia", "Balikpapan", 20, 3.6}

	fmt.Println(villa) // hasilnya apa? budi santosa, ...., 3.15
	fmt.Println(villa.ipk)
	fmt.Println("==================================")

	// Panggi function struct
	villa.halloMahasiswa("Jarwo")

}

// Deklarasi struct sebagai function
func (mhs mahasiswa) halloMahasiswa(name string) {
	fmt.Println("Halo Pak Dosen ", name, ". Nama saya adalah ", mhs.name)
}
