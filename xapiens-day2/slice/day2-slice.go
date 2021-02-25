package main

import "fmt"

func main() {
	// var _ = []string{"apel", "mangga"}      // ini array/slice? -slice
	// var _ = [2]string{"jeruk", "pisang"}    // ini array/slice? -array
	// var _ = [...]string{"anggur", "durian"} // ini array/slice? -array

	//## slice ambil value dari array

	var buah = [...]string{"apel", "mangga", "jeruk",
		"pisang", "anggur", "durian", "kiwi"}
	var newBuah = buah[0:] //print semua isi dari array buah
	var newBuah2 = buah[3:]
	// sahlan : pisang - selesai
	// dimas : apel - jeruk
	var newBuah3 = buah[:0]
	// dimas : apel - kiwi
	// sahlan : [] / apel
	var newBuah4 = buah[:3]
	var newBuah5 = buah[3:3]
	// dimas : pisang
	// sahlan : pisang
	// wawan : [ ]
	var newBuah6 = buah[1:3]
	// sahlan : mangga & jeruk
	// dimas : mangga & jeruk
	var newBuah7 = buah[:]
	//sahlan : [ ]
	// dimas : value dari buah

	fmt.Println(newBuah)
	fmt.Println(newBuah2)
	fmt.Println(newBuah3)
	fmt.Println(newBuah4)
	fmt.Println(newBuah5)
	fmt.Println(newBuah6)
	fmt.Println(newBuah7)

	// ## fungsi len, cap, append
	var buahBuahan = []string{"apel", "mangga", "jeruk", "pisang", "kiwi"}

	fmt.Println(buahBuahan)
	//fungsi len / length
	fmt.Println(len(buahBuahan)) //panjangnya? 5

	//fungsi cap / kapasitas
	fmt.Println(cap(buahBuahan)) //kapasitasya? 5

	var buahBuahan1 = buahBuahan[0:3]
	fmt.Println(buahBuahan1)
	fmt.Println(len(buahBuahan1))
	fmt.Println(cap(buahBuahan1))
	// buahBuahan  -> ["apel", "mangga", "jeruk", "pisang", "kiwi"]
	// buahBuahan1 -> ["apel", "mangga", "jeruk", -, -]

	var buahBuahan2 = buahBuahan[1:3]
	fmt.Println(buahBuahan2)
	fmt.Println(len(buahBuahan2))
	fmt.Println(cap(buahBuahan2))
	// buahBuahan  -> ["apel", "mangga", "jeruk", "pisang", "kiwi"]
	// buahBuahan2 -> ["mangga", "jeruk", -, -]
	fmt.Println("============================================")
	//append
	var fruits = []string{"apel", "mangga", "jeruk"}
	var aFruits = append(fruits, "kiwi")

	fmt.Println(fruits)
	fmt.Println(aFruits)
	// fmt.Println(fruits)
}
