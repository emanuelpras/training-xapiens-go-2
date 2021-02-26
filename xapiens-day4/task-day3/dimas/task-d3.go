package main

import "fmt"

func main() {
	fmt.Println("halo selamat datang di kalk sederhana")
	var a int
	menu()
	fmt.Scanln(&a)

	//dibungkus loop
	switch a {
	case 1:
		fmt.Println("kalkulator")
		functions := map[string]func(int, int) int{
			"+": tambah,
			"*": kali,
			"-": kurang,
			"/": bagi,
		}
		var currentNumber int
		fmt.Println("masukkan angka:")
		fmt.Scan(&currentNumber)

		for true {
			var functionName string
			var number int

			fmt.Println("(+ - * /) : ")
			fmt.Scan(&functionName)
			if functionName == "#" {
				main()
			}

			fmt.Println("Silahkan masukkan angka:")
			fmt.Scan(&number)
			currentNumber = functions[functionName](currentNumber, number)
			fmt.Println("Hasilnya:")
			fmt.Println(currentNumber)
		}
	case 2:
		fmt.Println("Menu History")
	case 3:
		fmt.Println("terimakasih sudah menggunakan kalulator ini")
	default:
		fmt.Println("Menu yang anda masukkan salah. Aplikasi akan keluar")
		fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
	}
}

func menu() {
	fmt.Println("1. Kalkulator")
	fmt.Println("2. history")
	fmt.Println("3. exit")
	fmt.Print("masukkan menu:")
}

func tambah(x, y int) int {
	return x + y
}
func kali(x, y int) int {
	return x * y
}
func kurang(x, y int) int {
	return x - y
}
func bagi(x, y int) int {
	return x / y
}
