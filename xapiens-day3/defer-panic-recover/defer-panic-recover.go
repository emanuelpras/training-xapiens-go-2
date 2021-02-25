// Defer --> try
// Panic --> error
// Recover --> catch

//Error terjadi saat
// 1. tipe data nggak sama
// 2. 10/0 atau 5/0

package main

import "fmt"

func pembagian(value int) {
	defer logRecover()
	total := 10 / value
	fmt.Println("Hasil pembagian : ", total)
}

func logRecover() {
	pesanError := recover()
	if pesanError != nil {
		fmt.Println("Ada error nih, pesan errornya : ", pesanError)
	}
}

func main() {
	pembagian(0)
}
