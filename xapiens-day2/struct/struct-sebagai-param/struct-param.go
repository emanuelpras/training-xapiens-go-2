package main

import "fmt"

type mahasiswa struct {
	nama string
	asal string
	age  int
	ipk  float64
}

// Deklarasi struct sebagai method
func (mhs mahasiswa) halloMahasiswa(name string) {
	fmt.Println("Hello ", name, "nama saya ", mhs.nama)
}

func main() {
	villa := mahasiswa{"Villa", "Banjarmasin", 19, 3.4}
	villa.halloMahasiswa("wawan ")
}
