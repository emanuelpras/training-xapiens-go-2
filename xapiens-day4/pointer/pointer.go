//analogi code

/*
	var kotaA, kotaB string
	kotaA = "Yogyakarta"
	kotaB = kotaA

	fmt.Println()
	outputnya kotaA = Yogyakarta
	outputnya kotaB = Yogyakarta

	kotaA = "Surabaya"

	fmt.Println
	output kotaA = Surabaya
	outputnya kotaB = Yogyakarta ???

*/

package main

import "fmt"

func main() {
	var kotaA, kotaB string
	kotaA = "Yogyakarta"
	kotaB = kotaA

	fmt.Println(kotaA)
	fmt.Println(kotaB)
	fmt.Println("=========================")

	kotaA = "Surabaya"
	kotaB = kotaA
	fmt.Println(kotaA)
	fmt.Println(kotaB)

	// Deklarasi pointer di variable
	var cityA string
	var cityB *string

	cityA = "Bandung"
	cityB = &cityA

	fmt.Println("sebelum diubah value cityA", cityA)
	fmt.Println("sebelum diubah memory cityA", &cityA)
	fmt.Println("sebelum diubah memory cityB", cityB)
	fmt.Println("sebelum diubah value cityB", *cityB)
	fmt.Println("=========================")

	cityA = "Padang"

	fmt.Println("setelah diubah value cityA", cityA)
	fmt.Println("setelah diubah memory cityA", &cityA)
	fmt.Println("setelah diubah memory cityB", cityB)
	fmt.Println("setelah diubah value cityB", *cityB)
	fmt.Println("=========================")

	*cityB = "Malang"
	fmt.Println("cityB diubah value cityA", cityA)
	fmt.Println("cityB diubah memory cityA", &cityA)
	fmt.Println("cityB diubah memory cityB", cityB)
	fmt.Println("cityB diubah value cityB", *cityB)
	fmt.Println("=========================")

	cityA = "Medan"
	fmt.Println(cityA)
	fmt.Println(cityB)
	fmt.Println(*cityB)
	fmt.Println("=========================")

	//Deklarasi pointer di sebuah struct
	fmt.Println("==========Struct=========")
	address1 := address{"Surabaya", "Jatim", "Indonesia"}
	var adress2 *address = &address1

	fmt.Println("Iterasi 1 : adress1 = ", address1)
	fmt.Println("Iterasi 1 : adress2 = ", adress2)
	fmt.Println("=========================")

	//sesi 2
	var address3 *address
	address3 = &address1
	// = &address1
	address3.city = "Jombang"
	fmt.Println(&address1)
	fmt.Println("Iterasi 2 : adress1 = ", &address1) //kira kira valuenya? -Dimas {Jombang Jatim Indonesia}
	fmt.Println("Iterasi 2 : adress2 = ", adress2)
	fmt.Println("Iterasi 2 : adress3 = ", address3) //kira kira valuenya? -Sahlan error, karena address tipe pointer
	fmt.Println("=========================")

	var address4 = &address{"Palangka", "Kalteng", "Indonesia"}
	fmt.Println("Iterasi 3 : adress1 = ", address1) //Jombang Jatim Indo
	fmt.Println("Iterasi 3 : adress2 = ", adress2)  //Palangka kalteng indo
	fmt.Println("Iterasi 3 : adress3 = ", address3) //jombang jatim indo
	fmt.Println("Iterasi 3 : adress4 = ", address4) //palangka kalteng indo --seperti buat baru
	fmt.Println("=========================")

	var address5 *address
	address5 = &address1
	address5.city = "Makassar"
	fmt.Println("Iterasi 4 : adress1 = ", address1) //berubah karena refer ke address 1
	fmt.Println("Iterasi 4 : adress2 = ", adress2)  //berubah karena refer ke address 1
	fmt.Println("Iterasi 4 : adress3 = ", address3) //berubah karena refer ke address 1
	fmt.Println("Iterasi 4 : adress4 = ", address4) //tetap
	fmt.Println("Iterasi 4 : adress5 = ", address5) //berubah
	fmt.Println("=========================")

	changeCity(adress2)
	fmt.Println("Iterasi 5 : adress1 = ", address1) //berubah karena refer ke address 1
	fmt.Println("Iterasi 5 : adress2 = ", adress2)  //berubah karena refer ke address 1
	fmt.Println("Iterasi 5 : adress3 = ", address3) //berubah karena refer ke address 1
	fmt.Println("Iterasi 5 : adress4 = ", address4) //tetap
	fmt.Println("Iterasi 5 : adress5 = ", address5) //berubah
	fmt.Println("=========================")

	//val1, val2, val3
	var value1, value2, value3 float64 = 5.0, 5.0, 10.0
	var history []string
	history = append(history, tambah(value1, value2, &value3))

	fmt.Println(history) // hasilnya apa?  5.0000 + 5.0000 = 10.0000
}

func tambah(val1, val2 float64, val3 *float64) (historySementara string) {
	*val3 = val1 + val2
	historySementara = historyKalkulasi(val1, val2, val3, " + ")
	return
}

func historyKalkulasi(val1, val2 float64, val3 *float64, jenis string) (historySementara string) {
	historySementara = fmt.Sprintf("%f", val1) + jenis + fmt.Sprintf("%f", val2)
	historySementara = historySementara + " = " + fmt.Sprintf("%f", *val3)
	return
}

func changeCity(add *address) {
	add.city = "Semarang"
}

type address struct {
	city     string
	province string
	country  string
}
