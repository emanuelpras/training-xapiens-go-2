package main

import (
	"fmt"
	"strings"
)

func main() {
	// ## operator aritmatika
	value := ((5 + 5) % 2)

	fmt.Println(value)

	// ## Operator perbandingan
	value1 := (((2+6)%3)*4 - 2) / 3 //ini hasilnya berapa? 2
	isEqual := (value1 == 2)

	fmt.Println(isEqual) //ini kalau saya run, hasilnya? true

	city := "Yogyakarta"
	filter := strings.Contains(city, "d")

	// fmt.Printf("%T", filter)
	fmt.Println(filter)
	fmt.Println("=======================")
	//## Operator logika
	left := false
	right := true

	leftAndRight := left && right
	fmt.Println(leftAndRight)

	leftOrRight := left || right
	fmt.Println(leftOrRight) // true

	leftReverse := !left
	fmt.Println(leftReverse) //true

	rightReverse := !right
	fmt.Println(rightReverse) // false
}
