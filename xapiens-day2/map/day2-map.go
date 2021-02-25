package main

import "fmt"

func main() {
	//## Map tanpa inisialisasi value di awal
	var age map[string]int
	age = map[string]int{}

	age["Emanuel"] = 17
	age["Prasetyawan"] = 27

	fmt.Println(age["Emanuel"])

	//## Map dengan inisialisasi value di awal
	var ages = map[string]int{
		"Emanuel":     17,
		"Prasetyawan": 27,
	}

	fmt.Println(ages["Prasetyawan"])
	delete(ages, "Prasetyawan")
	fmt.Println(ages["Prasetyawan"]) //hasilnya apa? kosong / error
	fmt.Println("============================")

	//Combine slice & array
	var mahasiswa = []map[string]string{
		map[string]string{"nama": "emanuel", "nim": "123"},
		map[string]string{"nama": "sahlan", "nim": "234"},
		map[string]string{"nama": "dimas", "nim": "345"},
	}

	fmt.Println(len(mahasiswa))
	fmt.Println(mahasiswa[1]["nim"])
}
