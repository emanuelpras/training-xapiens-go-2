package main

import (
	"fmt"
	"strconv"
)

func tambah(value1, value2 float32) float32 {
	return value1 + value2
}

func kurang(value1, value2 float32) float32 {
	return value1 - value2
}

func kali(value1, value2 float32) float32 {
	return value1 * value2
}

func bagi(value1, value2 float32) float32 {
	return value1 / value2
}

func menu() {
	fmt.Println("Pilih menu :")
	fmt.Println("1. History")
	fmt.Println("2. Kalkulator")
	fmt.Println("3. Exit")
	fmt.Println("Masukan pilihan menu")
}

func main() {
	var inputName, inputMenu, inputKeluarMenu, jenisOperator string
	var total, number float32
	fmt.Print("Masukan nama : ")
	fmt.Scanln(&inputName)

	for {
		menu()
		fmt.Scanln(&inputMenu)
		if inputMenu == "3" {
			break
		} else if inputMenu == "2" {
			state := 1
			for {
				if state%2 == 1 {
					fmt.Println("Masukan angka, untuk keluar masukan #")
					fmt.Println("Hasil", total)
					fmt.Scanln(&inputKeluarMenu)

					number, _ = strconv.Atoi(inputKeluarMenu)

					if state == 1 {
						total = tambah(float32(number), total)
					} else {
						jenisOperator = inputKeluarMenu
						fmt.Println("Masukan angka, untuk keluar masukan #")
						fmt.Println("Hasil", total)
						fmt.Scanln(&inputKeluarMenu)

						if jenisOperator == "+" {
							fmt.Println("HAHAHAHAHA")
						}
					}

					if inputKeluarMenu == "#" {
						break
					}
				} else {
					fmt.Println("Masukan operator aritmatika (+ - / *)")
				}

				state++
			}

			// fmt.Scanln(&inputAngka)
		}
	}
}
