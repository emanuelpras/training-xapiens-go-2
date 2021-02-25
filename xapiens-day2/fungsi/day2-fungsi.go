package main

import (
	"fmt"
)

func main() {
	//fungsi dengan param
	for i := 0; i < 5; i++ {
		funcB(i)
	}

	funcC() //fungsi tanpa param

	//func dengan return
	var value = funcBC()
	fmt.Println(value)
}

func funcBC() string {
	return "Prasetyawan"
}

//fungsi tanpa param
func funcC() {
	fmt.Println("Thanks")
}

//fungsi dengan param
func funcB(i int) {
	fmt.Println("Angka ", i)
}
