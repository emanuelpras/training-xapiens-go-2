package main

import (
	"errors"
	"fmt"
)

func logRecover() {
	pesanError := recover()
	if pesanError != nil {
		fmt.Println("Terjadi error karena ", pesanError)
	}
}

func pembagian(value int) {
	defer logRecover()
	total := 10 / value
	fmt.Println(total)
}

//fungsi kita memang pengen bikin error BISA DI SKIP
func pembagianDenganError(value int) (int, error) {
	if value == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return (10 / value), nil
	}

}

func main() {
	var value int
	fmt.Println("Masukan angka pembagi :")
	fmt.Scanln(&value)
	pembagian(value)
	fmt.Println("=============================")

	//BISA DI SKIP
	fmt.Println("Masukan angka pembagi :")

	result, err := pembagianDenganError(value)
	fmt.Println("Hasil pembagian : ", result, "pesan error", err)
}
