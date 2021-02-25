package main

import (
	"fmt"
)

func main() {
	// ## array tanpa inisialisasi value di awal
	var cities [3]string
	cities[0] = "Yogyakarta"
	cities[1] = "Aceh"
	cities[2] = "Jakarta"

	fmt.Println(cities)    // outputnya apa? [Yogya Aceh jakarta]
	fmt.Println(cities[1]) // outputnya apa? Aceh
	fmt.Println(len(cities))

	//## array dengan inisialisasi value di awal
	var cities2 = [3]string{"Surabaya", "Malang", "Depok"}

	fmt.Println(cities2)

	//## array pakai keyword make
	var buah = make([]string, 2)
	buah[0] = "apel"
	buah[1] = "mangga"

	fmt.Println(buah)

	// ## array tanpa define length
	var number = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(number)
	fmt.Println(len(number))

}
