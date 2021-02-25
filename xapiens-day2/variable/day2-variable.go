package main

import "fmt"

func main() {
	// ## variable dengan tipe data
	var fullName string = "Emanuel Dina Prasetyawan"
	var lastName string = "Prasetyawan"

	fmt.Println(fullName)
	fmt.Println(lastName)

	fullName = "Emanuel Prasetyawan"

	fmt.Println(fullName)

	var age uint8 = 17

	fmt.Println(age)

	fmt.Println("===============================")

	// ## variable tanpa tipe data
	fullNames := "Emanuel Dina"

	fmt.Println(fullNames)

	// ## variable yang tidak dipakai
	_ = "Wawan"
}
